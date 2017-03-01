package main

func addScripts() {
	logger.Trace.Println("addScripts()")
	weePlayDate.AddScript("post-script", `function createCircles(){ for(i=0;i<500;i++) {
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
		document.getElementById("mySVG").appendChild(myCircle); } }   
		var svgNS = "http://www.w3.org/2000/svg"; var margin = 30;
		function rand(min, max) { min = Math.ceil(min); max = Math.floor(max); return Math.floor(Math.random() * (max - min)) + min; }
		function getRandomColor() { var letters = 'CDEF'; var color = '#'; for (var i = 0; i < 3; i++ ) { color += letters[Math.floor(Math.random() * 4)]; } return color; }`)
	weePlayDate.AddScript("post-script", `function showHide(id) { var e = $('#'+id)[0]; if (e.style.display == 'block' || e.style.display == '' || !e.style.display) { e.style.display = 'none'; } else { e.style.display = 'block'; } }`)

	weePlayDate.AddScript("init-script", `$('a.categoryButton').hover(function () {$(this).animate({backgroundColor: '#b2d2d2'})},function () {$(this).animate({backgroundColor: '#d3ede8'})}  );`)
	weePlayDate.AddScript("init-script", `$('div.categoryBox').hover(over, out); `)
	weePlayDate.AddScript("init-script", `$('.datepicker').datepicker({ choose: 'dateOfBirth', dateFormat: '`+Date_Format+`' }); $('.ui-widget').addClass('ptButton');`)
	weePlayDate.AddScript("init-script", `function over() { var span = this.getElementsByTagName('span'); $(span[0]).animate({opacity: 0.3});	$(span[1]).animate({color: 'white'}); } `)
	weePlayDate.AddScript("init-script", `function out() { var span = this.getElementsByTagName('span'); $(span[0]).animate({opacity: 0.7}); $(span[1]).animate({color: '#444'}); } `)
	weePlayDate.AddScript("init-script", `createCircles();`)
	//weePlayDate.AddScript("init-script", `$('.letter').draggable({containment:'#board',cursor:'move',zIndex:3,revert: true }); $('#word').droppable({ accept:'.letter', drop: function(event, ui) { $(ui.draggable).clone().appendTo($('#word')).css('top','10px').css('left','0px').css('margin-left','-20px'); } }); `)

	weePlayDate.AddScript("init-home-script", `setInterval( 
		function() { 
			$.ajax({ 
				url: '/home', 
				type: 'AJAX', 
				headers: { 'ajaxProcessingHandler':'update' }, 
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
						ul.append( item ); });
					$.each(obj["conversations"], function(room, talk) { 
						var ul = $("#"+room+" div.discussion ul"); 
						ul.empty(); 
						$.each(talk, function(index, message) { 
							if (message['author']=='') { 
								item = $(document.createElement('li')).text( decodeURIComponent(message['message'].replace(/\+/g, ' ')) ); 
								item.attr("class", "push"); } else { 
								item = $(document.createElement('li')).text(message['author']+':'+decodeURIComponent(message['message'].replace(/\+/g, ' ')) ); 
								item.attr("class", "pull"); } 
							ul.append( item ).append( '<br/>' ); 
						}); 
						$("#"+room+" .dialogHeader").addClass("update");
					}); 
				},
				error: function(data, textStatus, jqXHR) { 
					console.log("button fail!"); 
				} 
			}); 
		}, 5000);	`)
	weePlayDate.AddScript("init-home-script", `$('#local').draggable({containment:'#lower',cursor:'move',zIndex:3 }); $('#local').resizable();`)
	weePlayDate.AddScript("home-script", `function enterRoom(roomName) {
		if ($('#'+roomName).length) { return; }
		body = $('body');
		body.append("<div id='"+roomName+"' class='ptFloatDialog'><a class='modalButton closeModal' title='Close'>X</a><a class='modalButton minimizeModal' title='Minimize'>_</a><a class='modalButton whosHere' title='Who'>?</a><h4 class='dialogHeader ptListItem'>"+roomName+"</h4><div class='whoseThere subFloatHeader hidden'><ul></ul></div><div class='discussion scroll'><ul></ul></div><div class='text'><a class='modalButton sendMessage' title='Send'>S</a><a class='modalButton addEmoji' title='Emoji'>:)</a><a class='modalButton gesture' title='Gesture'>*</a><a class='modalButton invitePerson' title='Invite'>+</a><a class='modalButton other' title='Other'>?</a><textarea class='ptExpand' cols='40' rows='4' name='message-"+roomName+"'/></div></div>");
		$('.ptFloatDialog').each(function(){ $( this ).position({my:"right bottom"}); });
		$('#'+roomName+' a.closeModal').attr('onclick', 'exitRoom("'+roomName+'")');
		$('#'+roomName+' a.minimizeModal').attr('onclick', 'minimize("'+roomName+'")'); 
		$('#'+roomName+' a.whosHere').attr('onclick', 'whoseThere("'+roomName+'")');
		$('#'+roomName+' a.sendMessage').attr("onclick","var li = $( '<li class=\"ptListItem push\">'+$('#"+roomName+" textarea').val()+'</li>'); $('#"+roomName+" .discussion ul').append(li); sendMessage('"+roomName+"',$('#"+roomName+" textarea').val());");
		$('#'+roomName+' a.addEmoji').attr("onclick","alert('Not yet implemented.')");
		$('#'+roomName+' a.gesture').attr("onclick","alert('Not yet implemented.')");
		$('#'+roomName+' a.invitePerson').attr("onclick","alert('Not yet implemented.')");
		sendMessage(roomName, $('#'+roomName+' .text textarea').val()); 
	}`)
	weePlayDate.AddScript("home-script", `function whoseThere(room) { if ($('#'+room).hasClass('extendFloatDialog')) { $('#'+room).removeClass('extendFloatDialog'); $('#'+room+' .subFloatHeader').addClass('hidden'); $('#'+room+' .discussion').css('height',''); } else { $('#'+room).addClass('extendFloatDialog'); $('#'+room+' .subFloatHeader').removeClass('hidden'); $('#'+room+' .discussion').css('height','275px'); $.ajax({url: '/home',type: 'AJAX', headers: { 'ajaxProcessingHandler':'whoseThere' }, dataType: 'html', data: { 'roomName':room }, success: function(data, textStatus, jqXHR) { var ul = $("#"+room+" .whoseThere ul"); var obj = JSON.parse(data); ul.empty(); $.each(obj, function(index, party) { user = decodeURIComponent(party[1].replace(/\+/g, ' ')); name = decodeURIComponent(party[0].replace(/\+/g, ' ')); item = $(document.createElement('li')).text(name); item.attr("class", "ptListItem"); item.attr("onclick", "openProfile('"+user+"')"); ul.append( item ); }); }, error: function(data, textStatus, jqXHR) { console.log("fail!"); } }); }	}	`)
	weePlayDate.AddScript("home-script", `function exitRoom(roomName) { $('#'+roomName).remove(); $.ajax({url: '/home',type: 'AJAX', headers: { 'ajaxProcessingHandler':'exitRoom' },	dataType: 'html', data: { 'roomName':roomName }, success: function(data, textStatus, jqXHR) {}, error: function(data, textStatus, jqXHR) { console.log("exit room fail!"); }	});	}	`)
	weePlayDate.AddScript("home-script", `function sendMessage(room,message) { $.ajax({url: '/home',type: 'AJAX', headers: { 'ajaxProcessingHandler':'message' },	dataType: 'html', data: { 'roomName':room,'message':encodeURIComponent(message) }, success: function(data, textStatus, jqXHR) { $("#"+room+" textarea").val(''); var ul = $("#"+room+" .discussion ul"); var obj = JSON.parse(data); ul.empty(); $.each(obj, function(index, message) { if (message['author']=='') { item = $(document.createElement('li')).text( decodeURIComponent(message['message'].replace(/\+/g, ' ')) ); item.attr("class", "push"); } else { item = $(document.createElement('li')).text( message['author']+':'+decodeURIComponent(message['message'].replace(/\+/g, ' ')) ); item.attr("class", "pull"); } ul.append( item ).append( '<br/>' ); }); ul.parent().scrollTop(ul.parent()[0].scrollHeight - ul.parent().height()); }, error: function(data, textStatus, jqXHR) { console.log("send message fail!"); $("#"+room+" textarea").val('') } }); }`)
	weePlayDate.AddScript("home-script", `function initiateRoom(roomName) { $.ajax({url: '/home',type: 'AJAX', headers: { 'ajaxProcessingHandler':'newRoom' },	dataType: 'html', data: { 'roomName':roomName, 'roomPass':'HaHa!' }, success: function(data, textStatus, jqXHR) { enterRoom(roomName); }, error: function(data, textStatus, jqXHR) { console.log("new room fail!"); }	});	}`)
	weePlayDate.AddScript("home-script", `function onShowProfileModal(user, name) { 
		$.ajax({
			url: '/home',
			type: 'AJAX', 
			headers: { 'ajaxProcessingHandler':'profile' },	
			dataType: 'html', 
			data: { 'user':user, 'name':name }, 
			success: function(data, textStatus, jqXHR) { 
				var obj = JSON.parse(data); 
				$("#personModal div .info").empty();
				$("#personModal div .info").append("<p>Name:"+obj["name"]+"</p>");
				$("#personModal div .info").append("<p>Age:"+obj["age"]+", "+obj["sex"]+"</p>");
				$("#personModal div .info").append("<p>"+obj["profile"]+"</p>");
				$("#personModal div #personProfilePic").attr('img','../img/test.jpg');
				var likes = obj["likes"].split("|");
				if (likes.length>0) {
					$("#personModal div .info").append("<ul>Likes:</ul>");
					$.each(likes, function(index, like) { $("#personModal div ul").append("<li>"+like+"</li>"); });
				}
			}, 
			error: function(data, textStatus, jqXHR) { console.log("onShowProfileModal fail!"); } }); }`)
	weePlayDate.AddScript("home-script", `function minimize(room) { 
		if ($('#'+room).hasClass('minimizedFloatDialog')) { 
			$('#'+room).removeClass('minimizedFloatDialog'); 
			$('#'+room+' .text').removeClass('hidden');  } 
		else { $('#'+room).addClass('minimizedFloatDialog'); 
			$('#'+room+' .text').addClass('hidden');
			$('#'+room).removeClass('extendFloatDialog');
			$('#'+room+' .subFloatHeader').addClass('hidden');
			$('#'+room).removeClass('extendFloatDialog');
		} }`)
	weePlayDate.AddScript("home-script", `function openArticle(articleId) { 
		$.ajax({url: '/home',type: 'AJAX', 
			headers: { 'ajaxProcessingHandler':'article' },	
			dataType: 'html', 
			data: { 'articleName':articleId }, 
			success: function(data, textStatus, jqXHR) { 
				var info = JSON.parse(data);
				var aModal = $('#articleModal div div');
				aModal.empty();
				aModal.append("<h2>"+info["title"]+"</h2><a class='ptButton' onclick=\"onShowProfileModal('"+info["user"]+"', '"+info["author"]+"'); $(location).attr('href','#personModal');\">"+info["author"]+"</a><img src='../img/"+info["pic"]+"'><p>"+info["text"]+"</p>");
				$(location).attr('href','#articleModal');
			}, 
			error: function(data, textStatus, jqXHR) { console.log("open article fail: "+textStatus); }	
		});
	}`)
	weePlayDate.AddScript("main-script", `function addKid() { var kid = $('#newKid'); var form = $('#family'); form.append("<input type='hidden' name='child"+child+"' value='"+kid.find('#newKid').val()+"|"+$.datepicker.formatDate("`+Date_Format+`", kid.find('#dob').datepicker('getDate'))+"|"+$('input[name=gender]:checked','#kid').val()+"'/>");
		form.find('ul').append("<li class='ptListItem'>"+kid.find('#newKid').val()+", "+$.datepicker.formatDate("`+Date_Format+`", kid.find('#dob').datepicker('getDate'))+", "+$('input[name=gender]:checked','#kid').val()+"<a class='edit' title='Close'>E</a></li>");
		kid[0].style.display = 'none'; kid.find('#newKid').val(''); $('input[name=gender]:checked','#kid').prop("checked",false); kid.find('#age').val('1'); child = child + 1; if (parent>0) $('#submitFamily')[0].style.display = 'inline'; }`)
	weePlayDate.AddScript("main-script", `function addParent() { var parentTag = $('#newParent'); var form = $('#family'); form.append("<input type='hidden' name='parent"+parent+"' value='"+parentTag.find('#newParent').val()+"|"+$('input[name=gender]:checked','#parent').val()+"'/>");
		form.find('ul').append("<li class='ptListItem'>"+parentTag.find('#newParent').val()+", "+$('input[name=gender]:checked','#parent').val()+"</li>");
		parentTag[0].style.display = 'none'; parentTag.find('#newParent').val(''); $('input[name=gender]:checked','#parent').prop("checked",false); parent = parent + 1; if (child>0) $('#submitFamily')[0].style.display = 'inline'; }`)

}
