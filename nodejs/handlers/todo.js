var Joi = require('joi');

var tododb = require('../db/tododb');

module.exports = {
    getAll: {
        handler: function(request, reply) {
            tododb.getAll(function(data) {
                if(data){
                    reply(data);
                } else {
                    reply({"success": false, "msg": "Error: Empty."});
                }
            });
        }
    },
    get: {
        handler: function(request, reply) {
            id = request.params.id;
            tododb.get(id, function(data) {
                if(data){
                    reply(data);
                } else {
                    reply({"success": false, "msg": "Error: Todo not found."});
                }
            });
        }, validate: {
            params: { id: Joi.number().integer().min(1) }
        }
    },
    create: {
        handler: function(request, reply) {
            var todo = {
                name: request.payload.name,
                done: request.payload.done
            };

            tododb.insert(todo, function(data) {
                reply({"success": true, "msg": "Todo created."});
            });

        }, validate: {
            payload: {
                name: Joi.string().required(),
                done: Joi.boolean().required()
            }
        }
    },
    update: {
        handler: function(request, reply) {
            id = request.params.id;
            var todo = {
                name: request.payload.name,
                done: request.payload.done
            };
            tododb.update(id, todo, function(data) {
                if(data){
                    reply({"success": true, "msg": "Todo updated."});
                } else {
                    reply({"success": false, "msg": "Error: Updating Todo."});
                }
            });
        }, validate: {
            params: { id: Joi.number().integer().min(1) },
            payload: {
                name: Joi.string().required(),
                done: Joi.boolean().required()
            }
        }
    },
    delete: {
        handler:function(request, reply) {
            id = request.params.id;
            tododb.delete(id, function(data) {
                if(data){
                    reply({"success": true, "msg": "Todo deleted."});
                } else {
                    reply({"success": false, "msg": "Error: Deleting Todo."});
                }
            });
        }, validate: {
            params: { id: Joi.number().integer().min(1) }
        }
    }
};
