var app = new Vue({
    el: '#app',
    data: {
        error: '',
        slides: [],
        groups: [],
        selectGroup: 'todos'
    },

    created: function() {
        Vue.http.interceptors.push((request, next) => {
            request.headers.set('x-access-token', localStorage.getItem('token'));    
            next();
        });
        
        this.fetchAllGroups();
        this.fetchAllSlides({});
    },

    watch: {
    	selectGroup: function() {
    		console.log(this.selectGroup);
    	}
    },

    methods: {

    	filtrar: function() {
    		var params = {params: {q: this.selectGroup}};

    		if (this.selectGroup === 'todos') {
    			params = {};
    		}

    		this.fetchAllSlides(params);
    	},

    	fetchAllGroups: function() {
            this.error = '';
            this.$http.get(baseURL('/groups')).then(function(response) {
                this.groups = response.body;
            }).catch(function(response) {
                this.error = response.body.message;
            });
        },

        fetchAllSlides: function(params) {
            this.error = '';
            this.$http.get(baseURL('/slides'), params).then(function(response) {
                this.slides = response.body;
            }).catch(function(response) {
                this.error = response.body.message;
            });
        },

        deleteSlide: function(key, slide) {

        	if (!confirm("Tem certeza que deseja remover esse card?")) {
        		return ;
        	}

            this.error = '';
            this.$http.delete(baseURL('/slides/' + slide.id)).then(function(response) {                
                this.slides.splice(key, 1);
            }).catch(function(response) {
                this.error = response.body.message;
            });
        }
    }
});