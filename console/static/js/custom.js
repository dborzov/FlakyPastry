$(document).ready(function(){
  console.log( "The document got loaded!" );

  dima_reacts = function(){
        var thecommand = $("#dima-query").val();
        window.theresponse = 'No response from server :('

        $.post('./query', thecommand,function(data) {
            console.log(data);
            window.theresponse = data;
            var view = {
             flaky_in: thecommand,
             flaky_out: window.theresponse
            };


            dima_paste = Mustache.render('<div class="row-fluid"><div class="span2">&nbsp;</div><div class="span8 flaky-query"><p class="lead text-center"> {{flaky_in}} </div><div class="span2">&nbsp;</div></div><div class="row-fluid"><div class="span2">&nbsp;</div><div class="span8 flaky-response"><p class="lead text-center"> {{flaky_out}} </div><div class="span2">&nbsp;</div></div>', view);

            $("#dimasscreen").prepend(dima_paste);
        });

    }

  $("#button-prompt").click(dima_reacts);

  $(document).keypress(function(e) {
    if(e.which == 13) {
        event.preventDefault();
        dima_reacts();
        $("#dima-query").val("");
    }
});



});