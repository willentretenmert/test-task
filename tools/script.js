import http from 'k6/http';

export const options = {
    ext: {
        loadimpact: {
            distribution: { 'amazon:us:ashburn': { loadZone: 'amazon:us:ashburn', percent: 100 } },
            apm: [],
        },
    },
    thresholds: {},
    scenarios: {
        Scenario_1: {
            executor: 'constant-arrival-rate',
            rate: 1000,
            timeUnit: '1s',
            duration: '900s',
            preAllocatedVUs: 10,
            maxVUs: 15,
            exec: 'default',
        },
    },
}

export default function () {
    for (let i = 100000; i < 999999; i++) {
        const url = `http://127.0.0.1:8111/bin?bin=${i}`;
        const res = http.get(url);
    }
}