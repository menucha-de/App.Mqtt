import { Component, OnInit, ViewChild, ElementRef, Input } from '@angular/core';
import { DataService } from '../data.service';
import { saveAs } from 'file-saver';
import { AccessControl } from '../model/accessControl';
import { UtilService } from '@peramic/shared';

@Component({
  selector: 'app-acc-control',
  templateUrl: './acc-control.component.html',
  styleUrls: ['./acc-control.component.scss']
})
export class AccControlComponent implements OnInit {
  @ViewChild('acUploader') fileDialogEl: ElementRef;
  @ViewChild('passUploader') fileDialogEl1: ElementRef;
  @Input() group: string;
  constructor(private data: DataService, private util: UtilService) { }
  iconFile = 'assets/images/file_new_grey.png';
  iconPass = 'assets/images/file_new_grey.png';
  btnDisabled = false;
  btnPassDisabled = false;
  allowAnonymous = true;
  ngOnInit() {
    this.checkFile();
    this.checkPassFile();
    this.getAccess();
  }
  getAccess() {
    this.data.getAccess().subscribe((data: AccessControl) => {
      this.allowAnonymous = data.anonymous;
    }, err => {
      this.util.showMessage('error', err.error);
    });
  }
  checkFile() {
    this.data.getAcFile().subscribe(() => {
      this.iconFile = 'assets/images/file_check.png';
      this.btnDisabled = false;
    }, () => {
      this.iconFile = 'assets/images/file_new_grey.png';
      this.btnDisabled = true;
    });
  }
  checkPassFile() {
    this.data.getPassFile().subscribe(() => {
      this.iconPass = 'assets/images/file_check.png';
      this.btnPassDisabled = false;
    }, () => {
      this.iconPass = 'assets/images/file_new_grey.png';
      this.btnPassDisabled = true;
    });
  }
  importPassFile(files: FileList) {
    this.util.showSpinner();
    this.data.uploadPassFile(files.item(0)).subscribe(() => {
      this.checkPassFile();
      this.util.hideSpinner();
    }, err => {
      this.util.showMessage('error', err.error);
      this.util.hideSpinner();
    });
    this.fileDialogEl1.nativeElement.value = null;
  }
  deletePassFile() {
    this.data.deletePassFile().subscribe(() => {
      this.checkPassFile();
      this.allowAnonymous = true;
    }, err => { this.util.showMessage('error', err.error); });
  }
  exportPassFile() {
    this.util.showSpinner();
    this.data.downloadPassFile().subscribe(file => {
      saveAs(file, 'PasswordFile');
      this.util.hideSpinner();
    }, (err) => {
      this.util.hideSpinner();
      this.util.showMessage('error', err.error);
    });
  }
  importAcFile(files: FileList) {
    this.util.showSpinner();
    this.data.uploadAcFile(files.item(0)).subscribe(() => {
      this.checkFile();
      this.util.hideSpinner();
    }, err => {
      this.util.showMessage('error', err.error);
      this.util.hideSpinner();
    }
    );
    this.fileDialogEl.nativeElement.value = null;
  }
  deleteAcFile() {
    this.data.deleteAcFile().subscribe(() => {
      this.checkFile();

    }, err => this.util.showMessage('error', err.error));
  }
  exportAcFile() {
    this.data.downloadAcFile().subscribe(file => {
      saveAs(file, 'AccessControlFile');
      this.util.hideSpinner();
    }, (err) => {
      this.util.hideSpinner();
      this.util.showMessage('error', err.error);
    });
  }
  updateAccess() {
    const ss: AccessControl = new AccessControl();
    ss.anonymous = this.allowAnonymous;
    if (!this.allowAnonymous) {
      if (this.btnPassDisabled) {
        this.util.showMessage('error', 'Please upload first a Password File');
        this.allowAnonymous = true;
        return;
      }
    }
    this.data.putAccess(ss).subscribe(() => {

    }, err => {
      this.util.showMessage('error', err.error);
      this.getAccess();
    });
  }
}
