import babylon from 'babylonjs';

window.addEventListener("load", function(evt) {
  var boids = {}

  var scene = new THREE.Scene();
  var camera = new THREE.PerspectiveCamera( 50, window.innerWidth/window.innerHeight, 0.1, 2000 );

  var renderer = new THREE.WebGLRenderer();
  renderer.setSize( window.innerWidth, window.innerHeight );
  document.body.appendChild( renderer.domElement );

  camera.position.z = 5;

  var geometry = new THREE.BoxGeometry( 1, 1, 1 );

  var ws;
  ws = new WebSocket("ws://localhost:3000/ws");
  ws.onopen = function(evt) {
    console.log("OPEN");
  }

  ws.onclose = function(evt) {
    console.log("CLOSE");
    ws = null;
  }

  ws.onerror = function(evt) {
    console.log("ERROR: " + evt.data);
  }

  ws.onmessage = function(evt) {
    boid = JSON.parse(evt.data);
    boids[boid.id] = boid;

    var material = new THREE.MeshBasicMaterial( { color: 0x00ff00 } );
    var cube = new THREE.Mesh( geometry, material );
    cube.position.x = boid.position[0];
    cube.position.y = boid.position[1];
    cube.position.z = 0;
    scene.add( cube );
  };

  var render = function () {
    requestAnimationFrame( render );

    renderer.render(scene, camera);
  };

  render();
});

function initializeThree() {

  return scene;
}
