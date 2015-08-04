var models = require('../models');

module.exports = {
    getAll: function(callback) {
        models.Todo.findAll().then(function(todos) {
            if(todos){
                callback(JSON.stringify(todos));
            } else {
                callback(false);
            }
        });
    },
    get: function(id, callback) {
        models.Todo.findById(id).then(function (todo) {;
            if(todo){
                callback(todo.get());
            } else {
                callback(false);
            }
        });
    },
    create: function(todo, callback) {
        models.Todo.create({
            name: todo.name,
            done: todo.done
        });
        callback(true);
    },
    update:function(id, todo, callback) {
        models.Todo.findById(id).then(function (t) {
            if(t){
                t.update({
                    name: todo.name,
                    done: todo.done
                });
                callback(true);
            } else {
                callback(false);
            }
        });
    },
    delete: function(id, callback) {
        models.Todo.findById(id).then(function (t) {
            if(t){
                t.destroy();
                callback(true);
            } else {
                callback(false);
            }
        });
    }
};
