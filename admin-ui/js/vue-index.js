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
                localStorage.setItem('token', response.body.token);
                localStorage.setItem('email', response.body.email);
                localStorage.setItem('id', response.body.id);
                
                window.location.href = '/manage.html';

            }).catch(function(response) {
                this.error = response.body.message;
            });
        }
    }
});