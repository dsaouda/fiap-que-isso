var $slick = null;

var app = new Vue({
    el: '#app',
    data: {
       cards: [],
       groups: [],
       selectGroup: '',
       checkboxEmbaralhar: false
    },

    created: function() {
        this.fetchAllGroups();        
    },

    watch: {
        selectGroup: function() {}
    },

    computed: {},
    updated: function() {

        if ($slick !== null) {
            $slick.slick('unslick');
            $slick = null;
        }

        if (this.cards.length > 0) {
            var html = $('#vue-template').html();

            $slick = $('.slider');    
            $slick.html(html);
            $slick.slick();
        }
    },

    methods: {

        fetchAllCards: function() {

            var params = this.selectGroup === 'todos' ? {params: {m: this.checkboxEmbaralhar } } : {
                params: {
                    q: this.selectGroup,
                    m: this.checkboxEmbaralhar
                }
            };    

            this.$http.get(baseUrlApp('/cards'), params).then(function(response) {
                this.cards = response.body;                
            });
        },

        fetchAllGroups: function() {
            this.$http.get(baseUrlApp('/groups')).then(function(response) {
                this.groups = response.body;                
            });
        }
        
    }
});