var Hapi = require('hapi');
var handlers = require('./handlers/todo');

var server = new Hapi.Server();
server.connection({port:8080});

server.route({
    method: 'GET',
    path: '/todo/',
    config: handlers.getAll
});

server.route({
    method: 'GET',
    path: '/todo/{id}',
    config: handlers.get
});

server.route({
    method: 'POST',
    path: '/todo/',
    config: handlers.create
});

server.route({
    method: 'PUT',
    path: '/todo/{id}',
    config: handlers.update
});

server.route({
    method: 'DELETE',
    path: '/todo/{id}',
    config: handlers.delete
});



server.start(function () {
    console.log('Server running at:', server.info.uri);
});
