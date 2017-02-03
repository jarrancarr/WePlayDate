package main

func addScripts() {
	logger.Trace.Println("addScripts()")
	weePlayDate.AddScript("pre-script", `$('a.categoryButton').hover(
		function () {$(this).animate({backgroundColor: '#b2d2d2'})},
		function () {$(this).animate({backgroundColor: '#d3ede8'})}  );`)
	weePlayDate.AddScript("pre-script", `$('div.categoryBox').hover(over, out); `)
	weePlayDate.AddScript("pre-script", `function over() {
		var span = this.getElementsByTagName('span');
		$(span[0]).animate({opacity: 0.3});
		$(span[1]).animate({color: 'white'}); } `)
	weePlayDate.AddScript("pre-script", `function out() {
		var span = this.getElementsByTagName('span');
		$(span[0]).animate({opacity: 0.7});
		$(span[1]).animate({color: '#444'}); } `)
	weePlayDate.AddScript("pre-script", `createCircles();`)
	weePlayDate.AddScript("pre-script",`
		$('.letter').draggable({
			containment:'#board',
			cursor:'move',
			zIndex:3,
			revert: true });
		$('#word').droppable({ 
			accept:'.letter', 
			drop: function(event, ui) { 
				$(ui.draggable).clone().appendTo($('#word')).css('top','10px').css('left','0px').css('margin-left','-20px');
			} 
		}); `)
	weePlayDate.AddScript("post-script", `
		function showHide(id) {
			var e = $('#'+id)[0];
			if (e.style.display == 'block' || e.style.display == '' || !e.style.display) {
				e.style.display = 'none';
			} else {
				e.style.display = 'block';
			}
		}`)		
	weePlayDate.AddScript("post-script", `
		var child = 0;
		function addKid() {
			var kid = $('#newKid');
			var form = $('#family');
			form.append("<input type='hidden' name='child"+child+"' value='"+kid.find('#newKid').val()+"|"+$('input[name=gender]:checked','#kid').val()+"|"+kid.find('#age').val()+"'/>");
			form.find('ul').append("<li class='ptListItem'>"+kid.find('#newKid').val()+", "+kid.find('#age').val()+", "+$('input[name=gender]:checked','#kid').val()+"</li>");
			kid[0].style.display = 'none';
			kid.find('#newKid').val('');
			$('input[name=gender]:checked','#kid').prop("checked",false);
			kid.find('#age').val('1');
			child = child + 1;
			if (parent>0) $('#submitFamily')[0].style.display = 'inline';
		}`)
	weePlayDate.AddScript("post-script", `
		var parent = 0;
		function addParent() {
			var parent = $('#newParent');
			var form = $('#family');
			form.append("<input type='hidden' name='parent"+parent+"' value='"+parent.find('#newParent').val()+"|"+$('input[name=gender]:checked','#parent').val()+"'/>");
			form.find('ul').append("<li class='ptListItem'>"+parent.find('#newParent').val()+", "+$('input[name=gender]:checked','#parent').val()+"</li>");
			parent[0].style.display = 'none';
			parent.find('#newKid').val('');
			$('input[name=gender]:checked','#parent').prop("checked",false);
			parent = parent + 1;
			if (child>0) $('#submitFamily')[0].style.display = 'inline';
		}`)
	weePlayDate.AddScript("post-script", `
		function createCircles(){
			for(i=0;i<500;i++) {
				var myCircle = document.createElementNS(svgNS,"circle");
				var color = getRandomColor();
				var height = rand(margin,screen.height*2);
				myCircle.setAttributeNS(null,"id","mycircle");
				myCircle.setAttributeNS(null,"cx",rand(margin,screen.width-margin));
				myCircle.setAttributeNS(null,"cy",height);
				myCircle.setAttributeNS(null,"r",rand(20,100-i/10));
				myCircle.setAttributeNS(null,"fill",color);
				myCircle.setAttributeNS(null,"stroke",color);
				myCircle.setAttributeNS(null,"stroke-width","1");
				myCircle.setAttributeNS(null,"fill-opacity","0.4");
				myCircle.setAttributeNS(null,"onclick","pop(this)");
				var animate = document.createElementNS(svgNS,"animate");
				animate.setAttributeNS(null,"attributeName","cy");
				animate.setAttributeNS(null,"from",height);
				animate.setAttributeNS(null,"to",height-screen.height-rand(margin,screen.height));
				animate.setAttributeNS(null,"dur","180s");
				animate.setAttributeNS(null,"begin","0s");
				animate.setAttributeNS(null,"repeatCount","indefinite");
				myCircle.appendChild(animate);		
				document.getElementById("mySVG").appendChild(myCircle);
			}
		}   
		var svgNS = "http://www.w3.org/2000/svg";  
		var margin = 30;


		function rand(min, max) {
		  min = Math.ceil(min);
		  max = Math.floor(max);
		  return Math.floor(Math.random() * (max - min)) + min;
		}
		
		function getRandomColor() {
		    var letters = 'CDEF';
		    var color = '#';
		    for (var i = 0; i < 3; i++ ) {
		        color += letters[Math.floor(Math.random() * 4)];
		    }
		    return color;
		}`)
	weePlayDate.AddScript("home-script", `
		function enterRoom(roomName) {
			if ($('#'+roomName).length) { return; }
			body = $('body');
			body.append("<div id='"+roomName+"' class='ptFloatDialog'><a class='closeModal' title='Close'>X</a><h4 class='dialogHeader'>"+roomName+"</h4><div class='scroll'><ul></ul></div><div><a class='sendMessage' title='Send'>S</a><a class='addEmoji' title='Emoji'>:)</a><a class='gesture' title='Gesture'>*</a><a class='invitePerson' title='Invite'>+</a><a class='other' title='Other'>?</a><textarea class='ptExpand' cols='40' rows='4' name='message-"+roomName+"'/></div></div>");
			$('.ptFloatDialog').each(function(){ $( this ).position({my:"right bottom"}); });
			$('#'+roomName+' a.closeModal').attr("onclick", "$('#"+roomName+"').remove();");
			$('#'+roomName+' a.sendMessage').attr("onclick","var li = $( '<li class=\"ptListItem push\">'+$('#"+roomName+" textarea').val()+'</li>'); $('#"+roomName+" ul').append(li); if ( $('#"+roomName+" ul li').length > 12) $('#"+roomName+" ul li:first').remove(); sendMessage('"+roomName+"',$('#"+roomName+" textarea').val());");
			$('#'+roomName+' a.addEmoji').attr("onclick","alert('Not yet implemented.')");
			$('#'+roomName+' a.gesture').attr("onclick","alert('Not yet implemented.')");
			$('#'+roomName+' a.invitePerson').attr("onclick","alert('Not yet implemented.')");
		}`)
	weePlayDate.AddScript("home-script", `
		function sendMessage(room,message) {
			// alert(room+": "+message);
			$.ajax({
					url: '/home',
					type: 'AJAX',
					headers: { 'ajaxProcessingHandler':'message' },
					dataType: 'html',
					data: { 'roomName':room,'message':message },
					success: function(data, textStatus, jqXHR) {
						$("#"+room+" textarea").val('')
					},
					error: function(data, textStatus, jqXHR) {
						console.log("send message fail!");
						$("#"+room+" textarea").val('')
					}
				});
		}`)
}