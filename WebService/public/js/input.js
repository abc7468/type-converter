(function($) {
    'use strict';
    $(function() {
        $('#inputBtn').on("click", function(event) {
            var body = $('#data').val();
            var data = {body: body};
            fetch("/send", {
                method: 'POST', // or 'PUT'
                body: data, // data can be `string` or {object}!
              })
              .catch(error => console.error('Error:', error))

    
        });
    })
    })(jQuery);

