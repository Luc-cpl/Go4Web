{{define "js"}}

	jQuery(document).ready(function(){
		jQuery('#send').submit(function(){
			var dados = jQuery( this ).serialize();

			jQuery.ajax({
				type: "POST",
				url: "/api/users/login",
				data: dados,
				success: function( data )
				{
					obj = JSON.parse(data);
					
                    if (obj.login == true){
                        window.location.replace("/");
                    }else{
						$('#alert').html(`<div class="alert alert-danger" role="alert">wrong email or password.
  											<a href="#" class="alert-link">Remember-me</a>
										  </div>`);
					}
				}
			});
			return false;
		});
	});

	jQuery(document).ready(function(){
		jQuery('#new').submit(function(){
			var dados = jQuery( this ).serialize();

			jQuery.ajax({
				type: "POST",
				url: "/api/users/new",
				data: dados,
				success: function( data )
				{
					obj = JSON.parse(data);
					
                    if (obj.login == true){
                       window.location.replace("/");
                    }else if (obj.err == "User already exist"){
						$('#err').html(`<div class="alert alert-danger" role="alert">`+ obj.err + `
  											<a href="#" class="alert-link">Remember-me</a>
										</div>`);

					}else{
						$('#err').html(`<div class="alert alert-danger" role="alert">`+ obj.err + `</div>`);
					}
				}
			});
			return false;
		});
	});

	function dropAlert(){
		$('#alert').html(``);
	}

{{end}}