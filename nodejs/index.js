var Hapi = require('hapi');
var routes = require('./routes/todo');
var models = require('./models');

var server = new Hapi.Server();
server.connection({port:8080});
server.route(routes);

models.sequelize.sync().then(function() {
    console.log('db sync');

    server.start(function () {
        console.log('Server running at:', server.info.uri);
    });
});
