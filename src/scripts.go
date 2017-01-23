package main

func addScripts() {
	logger.Trace.Println("addScripts()")
	wePlayDate.AddScript("pre-script", `$('a.categoryButton').hover(
		function () {$(this).animate({backgroundColor: '#b2d2d2'})},
		function () {$(this).animate({backgroundColor: '#d3ede8'})}  );`)
	wePlayDate.AddScript("pre-script", `$('div.categoryBox').hover(over, out); `)
	wePlayDate.AddScript("pre-script", `function over() {
		var span = this.getElementsByTagName('span');
		$(span[0]).animate({opacity: 0.3});
		$(span[1]).animate({color: 'white'}); } `)
	wePlayDate.AddScript("pre-script", `function out() {
		var span = this.getElementsByTagName('span');
		$(span[0]).animate({opacity: 0.7});
		$(span[1]).animate({color: '#444'}); } `)

	wePlayDate.AddScript("pre-script", `createCircles();`)
	
	wePlayDate.AddScript("post-script", `
		function showHide(id) {
			var e = $('#'+id)[0];
			if (e.style.display == 'block' || e.style.display == '' || !e.style.display) {
				e.style.display = 'none';
			} else {
				e.style.display = 'block';
			}
		}`)
		
	wePlayDate.AddScript("post-script", `
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
		
	wePlayDate.AddScript("post-script", `
		var parent = 0;
		function addParent() {
			var parent = $('#newParent');
			var form = $('#family');
			form.append("<input type='hidden' name='Parent"+parent+"' value='"+parent.find('#newParent').val()+"|"+$('input[name=gender]:checked','#parent').val()+"'/>");
			form.find('ul').append("<li class='ptListItem'>"+parent.find('#newParent').val()+", "+$('input[name=gender]:checked','#parent').val()+"</li>");
			parent[0].style.display = 'none';
			parent.find('#newKid').val('');
			$('input[name=gender]:checked','#parent').prop("checked",false);
			parent = parent + 1;
			if (child>0) $('#submitFamily')[0].style.display = 'inline';
		}`)
		
	wePlayDate.AddScript("post-script", `
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
}
