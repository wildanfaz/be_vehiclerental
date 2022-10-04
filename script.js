import http from 'k6/http';
import { check, sleep } from 'k6';
const token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWRtaW4iLCJyb2xlIjoiYWRtaW4iLCJleHAiOjE2NjQ4NTkxODh9.yb3uGoAXqBfQ310gN9rCn0203faobIA2T500yAxqa8U'

export const options = {
  stages: [
    { duration: '5s', target: 50 },
    { duration: '5s', target: 100 },
    { duration: '5s', target: 0 },
  ],
};

export default function () {
  const authHeaders = {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  };

  const res = http.get('https://fazdev-go-vehiclerental.herokuapp.com/api/v1/users', authHeaders);
  check(res, { 'status was 200': (r) => r.status == 200 });
  sleep(1);
}
