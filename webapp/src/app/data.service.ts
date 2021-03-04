import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Security, AccessControl } from './model/models';

@Injectable({
  providedIn: 'root'
})
export class DataService {
  private readonly baseUrl = 'rest/mqtt/';
  constructor(protected http: HttpClient) { }
  getAcFile() {
    return this.http.head(`${this.baseUrl}accesscontrol/acfile`);
  }
  uploadAcFile(file: File) {
    const headers = new HttpHeaders().append('Content-Type', 'application/octet-stream');
    return this.http.put<void>(`${this.baseUrl}accesscontrol/acfile`, file, { headers });

  }
  deleteAcFile() {
    return this.http.delete<void>(`${this.baseUrl}accesscontrol/acfile`);
  }
  downloadAcFile() {
    return this.http.get(`${this.baseUrl}accesscontrol/acfile`, {responseType: 'blob' });
  }
  getPassFile() {
    return this.http.head(`${this.baseUrl}accesscontrol/passwordfile`);
  }
  uploadPassFile(file: File) {
    const headers = new HttpHeaders().append('Content-Type', 'application/octet-stream');
    return this.http.put<void>(`${this.baseUrl}accesscontrol/passwordfile`, file, { headers });

  }
  deletePassFile() {
    return this.http.delete<void>(`${this.baseUrl}accesscontrol/passwordfile`);
  }
  downloadPassFile() {
    return this.http.get(`${this.baseUrl}accesscontrol/passwordfile`, {responseType: 'blob' });
  }
  getCaFile() {
    return this.http.head(`${this.baseUrl}security/trust`);
  }
  deleteCaFile() {
    return this.http.delete<void>(`${this.baseUrl}security/trust`);
  }
  uploadCaFile(file: File) {
    const headers = new HttpHeaders().append('Content-Type', 'application/octet-stream');
    return this.http.put<void>(`${this.baseUrl}security/trust`, file, { headers });

  }
  getServFile() {
    return this.http.head(`${this.baseUrl}security/keystore`);
  }
  deleteServFile() {
    return this.http.delete<void>(`${this.baseUrl}security/keystore`);
  }
  uploadServFile(file: File) {
    const headers = new HttpHeaders().append('Content-Type', 'application/octet-stream');
    return this.http.put<void>(`${this.baseUrl}security/keystore`, file, { headers });
  }
  getRevFile() {
    return this.http.head(`${this.baseUrl}security/revoclist`);
  }
  deleteRevFile() {
    return this.http.delete<void>(`${this.baseUrl}security/revoclist`);
  }
  uploadRevFile(file: File) {
    const headers = new HttpHeaders().append('Content-Type', 'application/octet-stream');
    return this.http.put<void>(`${this.baseUrl}security/revoclist`, file, { headers });

  }
  putPassPhrase(pass: string) {
    return this.http.put<void>(`${this.baseUrl}security/passphrase`, pass);
  }
  putSecurity(security: Security) {
    return this.http.put<void>(`${this.baseUrl}security`, security);
  }
  putAccess(access: AccessControl) {
    return this.http.put<void>(`${this.baseUrl}accesscontrol`, access);
  }
  getAccess() {
    return this.http.get(`${this.baseUrl}accesscontrol`);
  }
  getSecurity() {
    return this.http.get(`${this.baseUrl}security`);
  }
}
