import { Injectable } from '@angular/core';
import {environment} from "../../environments/environment";

@Injectable({
  providedIn: 'root'
})
export class ApiAppService {
  baseURL = environment.production ? "" : "http://127.0.0.1:5000";
  constructor() { }
}
