var Joi = require('joi');

module.exports = {
    getAll: {
        handler: function(request, reply) {
            reply('getAll');
        }
    },
    get: {
        handler: function(request, reply) {
            id = request.params.id;
            reply('get' + id);
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
            reply('create' + todo);
        }, validate: {
            payload: {
                name: Joi.string().required(),
                done: Joi.boolean().required()
            }
        }
    },
    update: {
        handler: function(request, reply) {
            var todo = {
                name: request.payload.name,
                done: request.payload.done
            };
            reply('update' + todo);
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
            reply('delete');
        }, validate: {
            params: { id: Joi.number().integer().min(1) }
        }
    },
};
