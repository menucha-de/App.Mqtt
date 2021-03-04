import { Component, OnInit, ViewChild, ElementRef, Input } from '@angular/core';
import { DataService } from '../data.service';

import { Security } from '../model/security';
import { UtilService } from '@peramic/shared';

@Component({
  selector: 'app-security',
  templateUrl: './security.component.html',
  styleUrls: ['./security.component.scss']
})
export class SecurityComponent implements OnInit {
  @ViewChild('uploaderCa') fileDialogEl: ElementRef;
  @ViewChild('uploaderServ') fileDialogEl1: ElementRef;
  @ViewChild('uploaderRev') fileDialogEl2: ElementRef;
  @Input() group: string;
  constructor(private data: DataService, private util: UtilService) { }

  iconCaFile = 'assets/images/file_new_grey.png';
  iconServFile = 'assets/images/file_new_grey.png';
  iconRevFile = 'assets/images/file_new_grey.png';
  btnCaDisabled = false;
  btnServDisabled = false;
  btnRevDisabled = false;
  pass = '';
  ssl = false;
  clientVerification = false;
  ngOnInit() {
    this.checkCaFile();
    this.checkServFile();
    this.checkRevFile();
    this.getsecurity();
  }
  getsecurity() {
    this.data.getSecurity().subscribe((data: Security) => {
      this.ssl = data.SSL;
      this.clientVerification = data.clientVerification;
    }, err => {
      this.util.showMessage('error', err.error);
    });
  }
  checkCaFile() {
    this.data.getCaFile().subscribe(() => {
      this.iconCaFile = 'assets/images/file_check.png';
      this.btnCaDisabled = false;
    }, () => {
      this.iconCaFile = 'assets/images/file_new_grey.png';
      this.btnCaDisabled = true;
    });
  }
  importCaFile(files: FileList) {
    this.util.showSpinner();
    this.data.uploadCaFile(files.item(0)).subscribe(() => {
      this.checkCaFile();
      this.util.hideSpinner();
    }, err => {
      this.util.showMessage('error', err.error);
      this.util.hideSpinner();
    });
    this.fileDialogEl.nativeElement.value = null;
  }
  deleteCaFile() {
    this.data.deleteCaFile().subscribe(() => {
      this.checkCaFile();
      this.ssl = false;
    }, err => this.util.showMessage('error', err.error));
  }
  checkServFile() {
    this.data.getServFile().subscribe(() => {
      this.iconServFile = 'assets/images/file_check.png';
      this.btnServDisabled = false;
    }, () => {
      this.iconServFile = 'assets/images/file_new_grey.png';
      this.btnServDisabled = true;
    });
  }
  onBlurMethod() {
    this.data.putPassPhrase(this.pass).subscribe(() => {
    }, err => {
      this.util.showMessage('error', err.error);
    });
  }
  importServFile(files: FileList) {
    this.util.showSpinner();
    this.data.uploadServFile(files.item(0)).subscribe(() => {
      this.checkServFile();
      this.util.hideSpinner();
      this.pass = '';
    }, err => {
      this.util.showMessage('error', err.error);
      this.pass = '';
      this.util.hideSpinner();
    });
    this.fileDialogEl1.nativeElement.value = null;
  }
  deleteServFile() {
    this.data.deleteServFile().subscribe(() => {
      this.checkServFile();
      this.ssl = false;
    }, err => this.util.showMessage('error', err.error));
  }
  checkRevFile() {
    this.data.getRevFile().subscribe(() => {
      this.iconRevFile = 'assets/images/file_check.png';
      this.btnRevDisabled = false;
    }, () => {
      this.iconRevFile = 'assets/images/file_new_grey.png';
      this.btnRevDisabled = true;
    });
  }
  importRevFile(files: FileList) {
    this.util.showSpinner();
    this.data.uploadRevFile(files.item(0)).subscribe(() => {
      this.checkRevFile();
      this.util.hideSpinner();
    }, err => {
      this.util.showMessage('error', err.error);
      this.util.hideSpinner();
    });
    this.fileDialogEl2.nativeElement.value = null;
  }
  deleteRevFile() {
    this.data.deleteRevFile().subscribe(() => {
      this.checkRevFile();
    }, err => this.util.showMessage('error', err.error));
  }
  updateSSL() {
    const ss: Security = new Security();
    ss.SSL = this.ssl;
    ss.clientVerification = this.clientVerification;
    this.data.putSecurity(ss).subscribe(() => {
    }, err => {
      this.util.showMessage('error', err.error);
      this.getsecurity();
    });

  }
}
