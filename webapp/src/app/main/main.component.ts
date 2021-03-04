import { Component, OnInit, Inject } from '@angular/core';


@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.scss']
})
export class MainComponent implements OnInit {
  title = 'MQTT';

  constructor(
    @Inject('SHOW_FRAME') private _showFrame: boolean,

  ) { }
  get showFrame() {
    return this._showFrame;
  }

  ngOnInit() { }

}