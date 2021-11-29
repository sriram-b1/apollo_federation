const { ApolloServer } = require('apollo-server');
const { ApolloGateway } = require("@apollo/gateway");

const gateway = new ApolloGateway({
    serviceList: [
        { name: 'metro', url: 'http://localhost:8001/query' },
        { name: 'covid', url: 'http://localhost:8002/query' },
        { name: 'route', url: 'http://localhost:8003/query'}
    ],
});

const server = new ApolloServer({
    gateway,

    subscriptions: false,
});

server.listen().then(({ url }) => {
    console.log(`🚀 Server ready at ${url}`);
});