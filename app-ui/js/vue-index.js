var $slick = null;

var app = new Vue({
    el: '#app',
    data: {
       cards: [],
       groups: [],
       selectGroup: ''
    },

    created: function() {
        this.fetchAllGroups();        
    },

    watch: {
        selectGroup: function() {
            this.fetchAllCards();            
        }
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
            this.$http.get(baseUrlApp('/cards'), {params: {q: this.selectGroup}}).then(function(response) {
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