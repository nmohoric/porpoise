
$.getJSON("/api/containers", function( json ) { 

    jQuery.each(json, function(){
        var color = "red";
        if(this.Ports){
            color = "green";
        }
        var $div = $('#' + this.Id);

        if ( $div.length){
            $div.css({"background-color": color});
        } else {
            var $maindiv = $('#containers');
            $maindiv.append( "<div class='col-md-6' id='" + this.Id + "'><pre style='background-color: " + color + "'>" + JSON.stringify(this, undefined, 2) + "</pre></div>" );
        }
    });
});
