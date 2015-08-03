var Sequelize = require('sequelize');

var db = new Sequelize('todos', '', '', {
    dialect: 'sqlite',
    storage: '../database/todo.sqlite'
});

//sequelize.sync();

// model
var Todo = db.define('Todos',{
    name: Sequelize.STRING,
    done: Sequelize.BOOLEAN
}, {
    timestamps: false
});

module.exports = {
    getAll: function() {
        Todo.findAll()
        .then(function(todos) {
            console.log(JSON.stringify(todos));
        });
    },
    get: function(id) {
        Todo.findById(id).then(function (todo) {
            console.log(todo.get());
        });
    },
    create: function() {
        /*Todo.create({
            name: 'node RESTful api',
            done: false
        });*/
        console.log('insert');
    },
    update:function() {
        console.log('update');
    },
    delete: function() {
        console.log('delete');
    }
};
