import axios from 'axios';

export class locationService {
    getLocation() {
        return axios.get("https://192.168.68.102:8082/locations");
    }
} 