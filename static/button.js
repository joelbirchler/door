(function() {
  var req;

  function onResponse() {
    if (req.readyState === XMLHttpRequest.DONE && req.status != 200) {
      console.log(req);
    }
  }

  function onClick() {
    req = new XMLHttpRequest();
    req.onreadystatechange = onResponse;
    req.open('PUT', '/push');
    req.send();
  }

  document.getElementById('button').addEventListener('click', onClick);
})();
