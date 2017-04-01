import BABYLON from 'babylonjs';

window.addEventListener("DOMContentLoaded", function(evt) {
  var boids = {}

  var canvas = document.getElementById('renderCanvas');
  var engine = new BABYLON.Engine(canvas, true);

	var createScene = function() {
			// create a basic BJS Scene object
			var scene = new BABYLON.Scene(engine);

			// create a FreeCamera, and set its position to (x:0, y:5, z:-10)
			var camera = new BABYLON.FreeCamera('camera1', new BABYLON.Vector3(0, 5, -500), scene);

			// target the camera to scene origin
			camera.setTarget(BABYLON.Vector3.Zero());

			// attach the camera to the canvas
			camera.attachControl(canvas, false);

			// create a basic light, aiming 0,1,0 - meaning, to the sky
			var light = new BABYLON.HemisphericLight('light1', new BABYLON.Vector3(0,1,0), scene);

			// return the created scene
			return scene;
	};

	var scene = createScene();

	engine.runRenderLoop(function() {
		scene.render();
	});

  window.addEventListener('resize', function() {
    engine.resize();
  });

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
    var boid = JSON.parse(evt.data);

    if (!(boid.id in boids)) {
      var cube = BABYLON.Mesh.CreateBox(boid.id.toString(), 2, scene);
      boid.cube = cube;
      boids[boid.id] = boid;
    }
    boids[boid.id].position = boid.position;
    boids[boid.id].cube.position.x = boid.position[0] - 250.0;
    boids[boid.id].cube.position.y = boid.position[1] - 250.0;
  };
});
