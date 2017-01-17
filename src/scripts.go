package main

func addScripts() {
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

	wePlayDate.AddScript("post-script", `function createCircles(){
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
}	`)
}
