$(document).ready(function(){
	
	$("#form_registro").on("submit", function(e){
		console.log("entro");
		e.preventDefault();
		 user_name = $("#user_name").val()

		$.ajax({
			type:"POST",
			url: "http://localhost:8000/validate/",
			data:{
				"user_name": user_name
			},
			success: function(data){
				result(data)
			}
		})

	})

	function result(data){
		obj = JSON.parse(data)
		console.log(obj);
	}

});