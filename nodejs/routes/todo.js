var handlers = require('../handlers/todo');

module.exports = [
    {
        method: 'GET',
        path: '/todo/',
        config: handlers.getAll
    },
    {
        method: 'GET',
        path: '/todo/{id}',
        config: handlers.get
    },
    {
        method: 'POST',
        path: '/todo/',
        config: handlers.create
    },
    {
        method: 'PUT',
        path: '/todo/{id}',
        config: handlers.update
    },
    {
        method: 'DELETE',
        path: '/todo/{id}',
        config: handlers.delete
    }
];
