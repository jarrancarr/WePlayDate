package main

func addScripts() {
	logger.Trace.Println("addScripts()")
	weePlayDate.AddScript("init-script", `$('a.categoryButton').hover(
		function () {$(this).animate({backgroundColor: '#b2d2d2'})},
		function () {$(this).animate({backgroundColor: '#d3ede8'})}  );`)
	weePlayDate.AddScript("init-script", `$('div.categoryBox').hover(over, out); `)
	weePlayDate.AddScript("init-script", `function over() {
		var span = this.getElementsByTagName('span');
		$(span[0]).animate({opacity: 0.3});
		$(span[1]).animate({color: 'white'}); } `)
	weePlayDate.AddScript("init-script", `function out() {
		var span = this.getElementsByTagName('span');
		$(span[0]).animate({opacity: 0.7});
		$(span[1]).animate({color: '#444'}); } `)
	weePlayDate.AddScript("init-script", `createCircles();`)
	weePlayDate.AddScript("init-home-script", `setInterval(
		function() {
			$.ajax({
				url: '/home',
				type: 'AJAX',
				headers: { 'ajaxProcessingHandler':'getRooms' },
				dataType: 'html',
				data: {},
				success: function(data, textStatus, jqXHR) {
					var ul = $( "<ul/>", {"class": "ptButton"}); 
					var obj = JSON.parse(data); 
					$("#roomList").empty(); 
					$("#roomList").append(ul); 
					$.each(obj["rooms"], function(val, i) { 
						item = $(document.createElement('button')).text( val + '  ' + i + ' occupance' ); 
						item.attr("class", "ptListItem"); 
						item.attr("onclick","enterRoom('"+val+"')"); 
						ul.append( item ); 
					});
					$.each(obj["conversations"], function(room, talk) {
						var ul = $("#"+room+" div.discussion ul");
						ul.empty();
						$.each(talk, function(index, message) { 
							if (message['author']=='') {
								item = $(document.createElement('li')).text( decodeURIComponent(message['message'].replace(/\+/g, ' ')) ); 
								item.attr("class", "push"); 
							} else {
								item = $(document.createElement('li')).text( message['author']+':'+decodeURIComponent(message['message'].replace(/\+/g, ' ')) );
								item.attr("class", "pull"); 
							}
							ul.append( item ).append( '<br/>' ); 
						});
					});
				},
				error: function(data, textStatus, jqXHR) {
					console.log("button fail!");
				}
			});
		}, 500);	`)
	weePlayDate.AddScript("init-script",`$('.letter').draggable({
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
	weePlayDate.AddScript("post-script", `function showHide(id) {
			var e = $('#'+id)[0];
			if (e.style.display == 'block' || e.style.display == '' || !e.style.display) {
				e.style.display = 'none';
			} else {
				e.style.display = 'block';
			}
		}`)		
	weePlayDate.AddScript("main-script", `var child = 0; function addKid() {
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
	weePlayDate.AddScript("main-script", `var parent = 0; function addParent() {
			var parentTag = $('#newParent');
			var form = $('#family');
			form.append("<input type='hidden' name='parent"+parent+"' value='"+parent.find('#newParent').val()+"|"+$('input[name=gender]:checked','#parent').val()+"'/>");
			form.find('ul').append("<li class='ptListItem'>"+parent.find('#newParent').val()+", "+$('input[name=gender]:checked','#parent').val()+"</li>");
			parentTag[0].style.display = 'none';
			parentTag.find('#newKid').val('');
			$('input[name=gender]:checked','#parent').prop("checked",false);
			parent = parent + 1;
			if (child>0) $('#submitFamily')[0].style.display = 'inline';
		}`)
	weePlayDate.AddScript("post-script", `function createCircles(){
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
	weePlayDate.AddScript("home-script", `function enterRoom(roomName) {
			if ($('#'+roomName).length) { return; }
			body = $('body');
			body.append("<div id='"+roomName+"' class='ptFloatDialog'><a class='closeModal' title='Close'>X</a><a class='whosHere' title='Who'>?</a><h4 class='dialogHeader ptListItem'>"+roomName+"</h4><div class='whoseThere subFloatHeader hidden'><ul></ul></div><div class='discussion scroll'><ul></ul></div><div><a class='sendMessage' title='Send'>S</a><a class='addEmoji' title='Emoji'>:)</a><a class='gesture' title='Gesture'>*</a><a class='invitePerson' title='Invite'>+</a><a class='other' title='Other'>?</a><textarea class='ptExpand' cols='40' rows='4' name='message-"+roomName+"'/></div></div>");
			$('.ptFloatDialog').each(function(){ $( this ).position({my:"right bottom"}); });
			$('#'+roomName+' a.closeModal').attr('onclick', 'exitRoom("'+roomName+'")');
			$('#'+roomName+' a.whosHere').attr('onclick', 'whoseThere("'+roomName+'")');
			$('#'+roomName+' a.sendMessage').attr("onclick","var li = $( '<li class=\"ptListItem push\">'+$('#"+roomName+" textarea').val()+'</li>'); $('#"+roomName+" ul').append(li); if ( $('#"+roomName+" ul li').length > 12) $('#"+roomName+" ul li:first').remove(); sendMessage('"+roomName+"',$('#"+roomName+" textarea').val());");
			$('#'+roomName+' a.addEmoji').attr("onclick","alert('Not yet implemented.')");
			$('#'+roomName+' a.gesture').attr("onclick","alert('Not yet implemented.')");
			$('#'+roomName+' a.invitePerson').attr("onclick","alert('Not yet implemented.')");
			sendMessage(roomName,'');
		}`)
	weePlayDate.AddScript("home-script", `function whoseThere(room) {
		if ($('#'+room).hasClass('extendFloatDialog')) {
			$('#'+room).removeClass('extendFloatDialog');
			$('#'+room+' .subFloatHeader').addClass('hidden');
			
		} else {
			$('#'+room).addClass('extendFloatDialog');
			$('#'+room+' .subFloatHeader').removeClass('hidden');
			$.ajax({url: '/home',type: 'AJAX',
				headers: { 'ajaxProcessingHandler':'whoseThere' },	dataType: 'html',
				data: { 'roomName':room },
				success: function(data, textStatus, jqXHR) {					
					var ul = $("#"+room+" .whoseThere ul");
					var obj = JSON.parse(data);
					ul.empty();
					$.each(obj, function(index, party) { 
						item = $(document.createElement('li')).text( decodeURIComponent(party.replace(/\+/g, ' ')) ); 
						item.attr("class", "ptListItem"); 
						ul.append( item ); 
					});
				},
				error: function(data, textStatus, jqXHR) {
					console.log("fail!");
	}	});	}	}	`)
	weePlayDate.AddScript("home-script", `function exitRoom(roomName) {
			$('#'+roomName).remove();
			$.ajax({url: '/home',type: 'AJAX',
				headers: { 'ajaxProcessingHandler':'exitRoom' },	dataType: 'html',
				data: { 'roomName':roomName },
				success: function(data, textStatus, jqXHR) {},
				error: function(data, textStatus, jqXHR) {
					console.log("exit room fail!");
	}	});	}	`)
	weePlayDate.AddScript("home-script", `function sendMessage(room,message) {
			$.ajax({url: '/home',type: 'AJAX',
				headers: { 'ajaxProcessingHandler':'message' },	dataType: 'html',
				data: { 'roomName':room,'message':encodeURIComponent(message) },
				success: function(data, textStatus, jqXHR) {
					$("#"+room+" textarea").val('');
					var ul = $("#"+room+" .discussion ul");
					var obj = JSON.parse(data);
					ul.empty();
					$.each(obj, function(index, message) { 
						if (message['author']=='') {
							item = $(document.createElement('li')).text( decodeURIComponent(message['message'].replace(/\+/g, ' ')) ); 
							item.attr("class", "push"); 
						} else {
							item = $(document.createElement('li')).text( message['author']+':'+decodeURIComponent(message['message'].replace(/\+/g, ' ')) );
							item.attr("class", "pull"); 
						}
						ul.append( item ).append( '<br/>' ); 
					});
					ul.parent().scrollTop(ul.parent()[0].scrollHeight - ul.parent().height());
				},
				error: function(data, textStatus, jqXHR) {
					console.log("send message fail!");
					$("#"+room+" textarea").val('')
	}	});	}`	)
}