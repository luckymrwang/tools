require('skyapm-nodejs').start({ serviceName: 'nodejs-demo-code-1352021-02-08T11:02:09', directServers: '10.48.51.135:31800' });
//import { Domain } from 'domain';

var express = require('express');
var domain = require('domain');
var path = require('path');
var favicon = require('serve-favicon');
var logger = require('log4js');
var cookieParser = require('cookie-parser');
var bodyParser = require('body-parser');
var nunjucks = require('nunjucks');
var httpProxy = require('http-proxy');
var _ = require('underscore');
var fs = require("fs");
var helmet = require('helmet');
var tls = require('tls');
//session
var session = require('express-session');

//constants
const Consts = require('../constants');

const context = require('./lib/Context').getCurrentContext();
//login state filter
const loginStateFilter = require('./lib/LoginStateFilter')();
var iconv = require('iconv-lite')

const allowList = [/^\/login-api\/[a-z0-9_\-\/.%]+/i, /^\/node-api\/[a-z0-9_\-\/.%]+/i];

const server = function (app, callback) {
  logger.configure({
    replaceConsole: true,
    levels: {
      "[all]": "WARN"  //ERROR
    },
    appenders: [
      { type: "console" }
    ]
  });

  //设置项目部署根目录
  var compress = require('compression');
  app.use(compress());
  //根目录
  global.DIR_NAME = __dirname;
  global.context = context;
  let __DEV__ = process.env.NODE_ENV === 'development' ? true : false;
  var ENV = 'dev';
  if (__DEV__) {
    ENV = 'dev';
  } else {
    ENV = 'prod';
  }
  app.locals.ENV = ENV;
  context.setEnv(ENV);
  logger.configure(context.getResource('log4js.json'));
  process.on('uncaughtException', function(err) {
    logger.getLogger("Global").error(err.stack|| err);
    setTimeout(function() {
      process.exit(1);
    }, 100);
  });

  let VIEW_PATH = __DEV__ ? Consts.DEV_VIEWS : Consts.DIST_VIEWS;
  //sever only
  // VIEW_PATH = path.join(Consts.SERVER, 'views');
  // if (__DEV__) {
  app.disable('view cache');
  // }
  app.enable('trust proxy');
  app.set('trust proxy', 1); // trust first proxy
  app.disable('x-powered-by');
  // view engine setup
  // app.set('views', VIEW_PATH);
  // app.set('view engine', 'jade');
  // app.engine('html', require('ejs').renderFile);
  // app.set('view engine', 'html');
  //禁用重协商
  tls.CLIENT_RENEG_LIMIT = 0;
  tls.CLIENT_RENEG_WINDOW = 0;
  //安全Helmet相关设置
  if (!__DEV__) {
    app.use(helmet.contentSecurityPolicy({
      directives: {
        defaultSrc: ["'self'"],
        scriptSrc: ["'self'", "'unsafe-eval'"],
        styleSrc: ["'self'", "'unsafe-inline'"],
        imgSrc: ["'self'", 'data:', "blob:"],
        connectSrc: ["'self'", 'wss://*:*', 'ws://*:*'],
        fontSrc: ["'self'", 'data:'],
        objectSrc: ["'self'"],
        frameSrc: ["'self'"],
        'frame-ancestors': ["'self'"]
      }
    }));
    app.use(helmet.xssFilter());
    app.use(helmet.frameguard({action: 'sameorigin'}));
    app.use(helmet.hsts({
      maxAge:31536000,
      includeSubDomains: true
    }));
    app.use(helmet.noSniff());
  }

  let nunjucksEnv = nunjucks.configure(VIEW_PATH, {
    autoescape: true,
    express: app,
    watch: __DEV__,
    noCache: __DEV__
  });
  // json formatting
  nunjucksEnv.addFilter('safeJson', function(obj) {
  });

  // app.set('view engine', 'nunjucks');

  // uncomment after placing your favicon in /public
  app.use(favicon(path.join(__dirname, '../static', 'favicon.ico')));
  // app.use(morganLogger('dev'));
  let httpLogger = logger.connectLogger(logger.getLogger("http"), {
    level: 'auto',
    format: (req, res, format) => {
      return format(`:remote-addr - ":method :url HTTP/:http-version" :status :content-length ":referrer" ":user-agent"`);
    }
  });
  app.use(httpLogger);
  app.use(cookieParser());
  let sessionConf = {
    name: Consts.IDENTITY_KEY,
    secret: 'chyingp',  // 用来对session id相关的cookie进行签名
    //store: new FileStore(),  // 本地存储session（文本文件，也可以选择其他store，比如redis的）
    saveUninitialized: true,  // 是否自动保存未初始化的会话，建议false
    resave: false,  // 是否每次都重新保存会话，建议false
    rolling: false,
    cookie: {
      secure: !__DEV__,
      // httpOnly: true,
      // domain: 'example.com',
      // path: 'foo/bar',
      maxAge: 24 * 60 * 60 * 1000  // 有效期，单位是毫秒
    }
  };
  if (!__DEV__) {
    var serviceObj = context.getResource('serviceAddr.json');
    var RedisStore = require('connect-redis')(session)
    sessionConf.store = new RedisStore({
      host:serviceObj.redis.host,
      port:serviceObj.redis.port,
      pass:serviceObj.redis.password,
      db:0
    });
  }
  app.use(session(sessionConf));
  app.use(loginStateFilter);
  callback&&callback();

  app.use(express.static(VIEW_PATH))
  
  app.use(bodyParser.json({
    verify:function(req, res, buf, encoding) {
      let str = typeof buf !== 'string' && encoding !== null
        ? iconv.decode(buf, encoding)
        : buf;
      try {
        JSON.parse(str)
      } catch (error) {
        res.status(400);
        res.json({
          errCode: "",
          errMessage: "数据错误",
          exceptionMsg: null,
          flag: false,
          resData: {}
        });
      }
    }
  }));
  
  app.use(bodyParser.urlencoded({ extended: false }));
  // csrf filter
  app.use('/', function(req, res, next) {
    // if (req.url.indexOf("login.html") > -1&&req.session.token) {
    //   res.redirect("/index.html");
    //   next()
    // }
    let referer = req.headers.referer || '';
    let method = req.method.toLowerCase();
    const validateMethods = ['get', 'post', 'put', 'delete', 'patch'];
    if (req.protocol == "http") {
      sessionConf.cookie.secure = false;
    }
    let loginUrl = req.baseUrl ? (req.baseUrl + '/login.html') : '/login.html';
    let isCAS = context.getResource('serviceAddr.json').isCAS||false;
    if (isCAS) {
      loginUrl = req.baseUrl ? (req.baseUrl + '/loginWithCas.html') : '/loginWithCas.html';
    }
    if (validateMethods.includes(method)) { // 如果是get、post、put、delete的一种必须验证host、referer、origin
      let protocol = req.protocol;
      let host = req.hostname;
      let originalUrl = protocol + "://" + host;
      const csrfHost = req.session['csrfHost'];
      if (req.url.indexOf('/loginWithCas.html')>=0 || req.url.indexOf('/loginWithToken.html')>=0 ) {
        // 第三方就放行
        req.session['csrfHost'] = 'can-redirect-request';
        next();
      } else if (allowList.some(item => item.test(req.url)) && method === 'get') {
        // get请求白名单放行
        next();
      } else if (referer.indexOf(originalUrl) === 0 && (originalUrl === csrfHost || csrfHost === 'can-redirect-request')) {
        // 登录前要保存csrfHost 如果是第三方登录就给host设置为can-redirect-request host合法放行
        next();
      } else {
        const csrfHost = req.session.csrfHost;
        if (!csrfHost) {
          if (req.url.startsWith('/api')) {
            res.status(200);
            res.json({flag: false, errCode: "SAFE_FAIL"});
          } else if (req.url.startsWith('/node-api')) {
            next()
          } else {
            res.redirect('/timeout.html');
          }
          return;
        }
        res.json({
          code: '400001',
          errMessage: 'invalid request'
        })
      }
    } else {
      next();
    }
  });
  app.use(function(req, res, next) {
    var reqdomain = domain.create();
    reqdomain.on("error", function(err) {
      logger.getLogger("Global").error(err.stack|| err);
      res.send(500, err.stack)
    })
    reqdomain.run(next)
  })
  //临时中间件,迎合前端使用
  app.use(/^.+\.(jsp|html)$/, function (req, res, next) {
    req.method = "GET";  //GET必须大写！！！no reason
    next();
  });
  

  /**
   *  initialize proxy
   */
  var proxyServer = httpProxy.createProxyServer();
  context.setResource('proxy', proxyServer);
  proxyServer.on("error", function(e, req, res) {
    logger.getLogger("proxy").error(e.message);
    res.end('something sent wrong')
  });
  //给健康检查增加接口
  app.get('/health', function(req, res, next) {
    res.json({state:'up'})
  });
  //增加region及X-Auth-Keep-Alive传参
  app.use('/', function(req, res, next) {
    if (req.url.startsWith('/api/')
      ||req.url.startsWith('/customApi/')
      ||req.url.startsWith('/node-api/')) {
      if (!req.headers.regionid) {
        req.headers.region = req.session.regionId||"";
      } else {
        req.headers.region = req.headers.regionid;
      }
      req.headers['X-Auth-Keep-Alive'] = !req.headers.polling;
      // req.session._garbage = Date.now();
      // req.session.touch();
      next();
    } else {
      next();
    }
  });
  //路由挂载
  var routerFactory = require('./RouterFactory');
  routerFactory.mount(context.getResource('routes.json'), app, context);
  //未匹配的路由重定向到首页
  app.use('/', function(req, res, next) {
    res.redirect(req.session.token?'/index.html':'/login.html');
    return;
  });
  // catch 404 and forward to error handler
  app.use(function (req, res, next) {
    var err = new Error('Not Found');
    err.status = 404;
    next(err);
  });

  // error handler
  app.use(function (err, req, res) {
    // set locals, only providing error in development
    res.locals.message = err.message;
    res.locals.error = req.app.get('env') === 'development' ? err : {};
    // render the error page
    // res.status(err.status || 500);
    let status = err.status || 500,
      result = {
        message: err.message,
        error: {},
        status: err.status || 500,
        ret: false
      };

    if (/\.json$/.test(req.path)) {
      res.status(status).send(result);
    } else {
      res.status(status).render('error', result);
    }
  });

  //监视serviceAddr.json文件变化
  fs.watch(path.join(__dirname, 'resources/default/serviceAddr.json'), function() {
    context.clearResource();
    context.getResource('serviceAddr.json');
  })
};

module.exports = server;
