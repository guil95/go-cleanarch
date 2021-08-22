db = db.getSiblingDB('cleanarch_go_db');

db.createCollection('users');

db.users.insertMany([
    {
        identifier: '28980f19-1a0e-4916-9520-0c37eb735fcf',
        name: 'Guilherme',
        lastname: 'Rodrigues',
        age: 26,
    }
]);