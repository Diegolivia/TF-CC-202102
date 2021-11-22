import { Component, OnInit } from '@angular/core';

interface User {
  value: string;
  viewValue: string;
}


export interface PeriodicElement {
  position: number  ;
  name: string;
  user_type : string;
  state : string;
}

const ELEMENT_DATA: PeriodicElement[] = [
  {position: 1, name: 'Richard', user_type: 'Importador', state: 'No Enviado'},
  {position: 2, name: 'Diego', user_type: 'Aduanero', state: ' Enviado'},
  {position: 3, name: 'Kevin', user_type: 'Aduanero', state: 'No Enviado'},
  {position: 4, name: 'Fernando', user_type: 'Importador', state: ' Enviado'},
  {position: 5, name: 'Piero', user_type: 'Aduanero', state: ' Enviado'},
  {position: 6, name: 'Juan Pablo', user_type: 'Importador', state: ' Enviado'},
  {position: 7, name: 'Baker', user_type: 'Proveedor', state: ' Enviado'},
  {position: 8, name: 'Jhosep', user_type: 'Proveedor', state: ' Enviado'},
  {position: 9, name: 'Oscar', user_type: 'Proveedor', state: 'No Enviado'},
  {position: 10, name: 'J', user_type: 'Aduanero', state : 'No Enviado'},
];
@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  displayedColumns: string[] = ['position', 'name', 'weight', 'symbol'];
  dataSource = ELEMENT_DATA;

  users: User[] = [
    {value: 'importador-0', viewValue: 'Importador'},
    {value: 'aduanero-1', viewValue: 'Aduanero'},
    {value: 'proveedor-2', viewValue: 'Proveedor'},
  ];
  ngOnInit(): void {
  }

}
