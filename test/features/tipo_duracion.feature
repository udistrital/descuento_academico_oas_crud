Feature: Validate API responses
  DESCUENTO_ACADEMICO_CRUD
  probe JSON reponses


Scenario Outline: To probe route code response  /tipo_duracion
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      

  Examples: 
  |method|route            |bodyreq               |codres       |
  |GET   |/v1/tipo_duracion|./files/req/Vacio.json|200 OK       |
  |GET   |/v1/tipo_duracio |./files/req/Vacio.json|404 Not Found|
  |POST  |/v1/tipo_duracio |./files/req/Vacio.json|404 Not Found|
  |PUT   |/v1/tipo_duracio |./files/req/Vacio.json|404 Not Found|
  |DELETE|/v1/tipo_duracio |./files/req/Vacio.json|404 Not Found|


Scenario Outline: To probe response route /tipo_duracion       
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      
  And the response should match json "<bodyres>"

  Examples:
  |method|route            |bodyreq               |codres         |bodyres                |
  |GET   |/v1/tipo_duracion|./files/req/Vacio.json|200 OK         |./files/res7/Vok1.json |
  |POST  |/v1/tipo_duracion|./files/req/Yt1.json  |201 Created    |./files/res7/Vok2.json |
  |POST  |/v1/tipo_duracion|./files/req/Vacio.json|400 Bad Request|./files/res7/Ierr1.json|
  |POST  |/v1/tipo_duracion|./files/req/Nt1.json  |400 Bad Request|./files/res7/Ierr2.json|
  |POST  |/v1/tipo_duracion|./files/req/Nt2.json  |400 Bad Request|./files/res7/Ierr3.json|
  |POST  |/v1/tipo_duracion|./files/req/Nt3.json  |400 Bad Request|./files/res7/Ierr4.json|
  |POST  |/v1/tipo_duracion|./files/req/Nt4.json  |400 Bad Request|./files/res7/Ierr5.json|
  |POST  |/v1/tipo_duracion|./files/req/Nt5.json  |400 Bad Request|./files/res7/Ierr6.json| 
  |POST  |/v1/tipo_duracion|./files/req/Nt6.json  |400 Bad Request|./files/res7/Ierr7.json| 
  |POST  |/v1/tipo_duracion|./files/req/Yt2.json  |400 Bad Request|./files/res7/Ierr8.json| 
  |PUT   |/v1/tipo_duracion|./files/req/Yt2.json  |200 OK         |./files/res7/Vok2.json |
  |GETID |/v1/tipo_duracion|./files/req/Vacio.json|200 OK         |./files/res7/Vok2.json |
  |DELETE|/v1/tipo_duracion|./files/req/Vacio.json|200 OK         |./files/res7/Ino.json  |
  |DELETE|/v1/tipo_duracion|./files/req/Vacio.json|404 Not Found  |./files/res7/Ierr9.json|
