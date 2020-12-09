// 모듈추가
const express = require('express');
const app = express();
var bodyParser = require('body-parser');
// 하이퍼레저 모듈추가+연결속성파일로드
const { FileSystemWallet, Gateway } = require('fabric-network');
const fs = require('fs');
const path = require('path');
const { send } = require('./sdk');
const { resolve } = require('path');
const ccpPath = path.resolve(__dirname, '..', 'network' ,'connection.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);
// 서버속성
const PORT = 8080;
const HOST = '0.0.0.0';
// app.use
app.use(express.static(path.join(__dirname, 'views')));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));
// 라우팅
// 1. 페이지라우팅
app.get('/', (req, res)=>{
    res.sendFile(__dirname + '/index.html');
})
app.get('/add-user', (req, res)=>{
    res.sendFile(__dirname + '/add-user.html');
})
app.get('/attend', (req, res)=>{
    res.sendFile(__dirname + '/attend.html');
})
app.get('/sit', (req, res)=>{
    res.sendFile(__dirname + '/sit.html');
})
app.get('/download', (req, res)=>{
    res.sendFile(__dirname + '/download.html');
})
app.get('/queryAnswer', (req, res)=>{
    res.sendFile(__dirname + '/queryAnswer.html');
})
app.get('/saveNote', (req, res)=>{
    res.sendFile(__dirname + '/saveNote.html');
})
app.get('/exit', (req, res)=>{
    res.sendFile(__dirname + '/exit.html');
})
app.get('/queryUser', (req, res)=>{
    res.sendFile(__dirname + '/queryUser.html');
})

// 2. REST라우팅
app.post('/user', async(req, res)=>{
    const {add_id, total_count} = req.body;
    let args=[add_id, total_count];
    send(1,"addStudent", args,res);
})
app.post('/user/attend', async(req, res)=>{
    const {enter_id, enter_class} = req.body;
    let args=[enter_id, enter_class]
    send(1,"attand", args ,res);
})
app.post('/user/sit', async(req, res)=>{
    const {sit_id, sit_class, sitNo} = req.body;
    let args=[sit_id, sit_class, sitNo]
    send(1,"sit", args,res);
})
app.post('/user/download', async(req, res)=>{
    const {download_id, download_class} = req.body;
    send(1,"download_material", [download_id, download_class],res);
})
app.post('/user/queryAnswer', async(req, res)=>{
    const {queryAnswer_id, query_class, query, answer} = req.body;
    send(1,"query_answer", [queryAnswer_id, query_class, query, answer],res);
})
app.post('/user/saveNote', async(req, res)=>{
    const {saveNote_id,saveNote_class, note} = req.body;
    send(1,"save_note", [saveNote_id, saveNote_class, note],res);
})
app.post('/user/exit', async(req, res)=>{
    const {exit_id,exit_class} = req.body;
    send(1,"exit", [exit_id, exit_class],res);
})
app.get('/user', async(req, res)=>{
    const {query_id} = req.query;
    let args=[query_id];
    send(0,"queryStudent", args,res);
})

// 서버시작
app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);