import http from 'k6/http';
import { check, sleep } from 'k6';
const token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYSIsInJvbGUiOiJhZG1pbiIsImV4cCI6MTY2NDcxNjQ3MX0.H_d-RAv0JcDBtPpnJEUhRzRR7vFTJ1sH3C-4LcgSkJA'

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
