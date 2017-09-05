var app = new Vue({
    el: '#app',
    data: {
        error: '',
        slides: []
    },

    created: function() {
        Vue.http.interceptors.push((request, next) => {
            request.headers.set('x-access-token', localStorage.getItem('token'));    
            next();
        });
        
        this.fetchAllSlides();
    },

    methods: {
        fetchAllSlides: function() {
            this.error = '';
            this.$http.get(baseURL('/slides')).then(function(response) {
                this.slides = response.body;
            }).catch(function(response) {
                this.error = response.body.message;
            });
        }
    }
});