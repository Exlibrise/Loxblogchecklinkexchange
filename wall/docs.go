/*
Package wall contains all paywall-related code.

This is the package you need to use for creating a middleware for one of the supported handlers / routers / frameworks.
For creating a middleware you only need to call one of the provided factory functions,
but all functions require a storage client (an implementation of the wall.StorageClient interface) as parameter.
You can either pick on