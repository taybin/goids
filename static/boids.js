window.addEventListener("load", function(evt) {
  var ws;
  ws = new WebSocket("ws://localhost:3000/ws");
  ws.onopen = function(evt) {
    console.log("OPEN");
    ws.send("hello");
  }
  ws.onclose = function(evt) {
    console.log("CLOSE");
    ws = null;
  }
  ws.onmessage = function(evt) {
    console.log("RESPONSE: " + evt.data);
  }
  ws.onerror = function(evt) {
    console.log("ERROR: " + evt.data);
  }

  initializeThree();
});

function initializeThree() {
  var scene = new THREE.Scene();
  var camera = new THREE.PerspectiveCamera( 75, window.innerWidth/window.innerHeight, 0.1, 1000 );

  var renderer = new THREE.WebGLRenderer();
  renderer.setSize( window.innerWidth, window.innerHeight );
  document.body.appendChild( renderer.domElement );

  var geometry = new THREE.BoxGeometry( 1, 1, 1 );
  var material = new THREE.MeshBasicMaterial( { color: 0x00ff00 } );
  var cube = new THREE.Mesh( geometry, material );
  scene.add( cube );

  camera.position.z = 5;

  var render = function () {
    requestAnimationFrame( render );

    cube.rotation.x += 0.1;
    cube.rotation.y += 0.1;

    renderer.render(scene, camera);
  };

  render();
}
