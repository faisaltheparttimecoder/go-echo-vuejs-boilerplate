let qs = require('qs')

export default {
    data: function () {
        return {
            endpoints: {
                backendapi: "/backendapi",
            },
        }
    },
    methods: {
        get: function (url) {  // Get Request
            return this.axios.get(url)
                .then((response) => Promise.resolve(response))
                .catch((error) => Promise.reject(error))
        },
        post: function (url, data) {  // Post Request
            return this.axios.post(url, qs.parse(data))
                .then((response) => Promise.resolve(response))
                .catch((error) => Promise.reject(error))
        },
        patch: function (url, data) {  // Update Request
            return this.axios.put(url, data)
                .then((response) => Promise.resolve(response))
                .catch((error) => Promise.reject(error))
        },
        delete: function (url) {  // Delete Request
            return this.axios.delete(url)
                .then((response) => Promise.resolve(response))
                .catch((error) => Promise.reject(error))
        }
    }
}