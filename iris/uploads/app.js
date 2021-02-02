require('skyapm-nodejs').start({ serviceName: 'nodejs-demo-code-1432021-02-02T15:09:20', directServers: '10.48.51.135:21594' });
const express = require('express')
const app = express()
const port = 3000

app.get('/', (req, res) => {
    res.send('Hello World!')
})

app.listen(port, () => {
    console.log(`Example app listening at http://localhost:${port}`)
})
