module.exports = {
    getAll: {
        handler: function(request, reply) {
            reply('getAll');
        }
    },
    get: {
        handler: function(request, reply) {
            reply('get');
        }
    },
    create: {
        handler: function(request, reply) {
            reply('create');
        }
    },
    update: {
        handler: function(request, reply) {
            reply('update');
        }
    },
    delete: {
        handler:function(request, reply) {
            reply('delete');
        }
    },
};
