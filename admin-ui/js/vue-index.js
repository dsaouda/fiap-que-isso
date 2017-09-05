var app = new Vue({
    el: '#app',
    data: {
        error: '',
        form: {
            email: '',
            password: ''
        }
    },

    methods: {
        login: function() {
            this.error = '';
            this.$http.post(baseURL('/login'), this.form).then(function(response) {
                            
                sessionStorage.setItem('token', response.body.token);
                sessionStorage.setItem('email', response.body.email);
                sessionStorage.setItem('id', response.body.id);
                window.location.href = '/manage.html';

            }).catch(function(response) {
                this.error = response.body.message;
            });
        }
    }
});