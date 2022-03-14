export default [{
    jwt: require("jsonwebtoken"),
    METABASE_SITE_URL: "http://172.16.10.204:3000",
    METABASE_SECRET_KEY: "17f65a954631b5e001a38fca68a8246f425403551aed0de36fdeafdf44dbd04f",
    src: "http://172.16.10.204:3000/embed/dashboard/1-portfolio-overview/#bordered=true&titled=true",
    payload: {
        resource: { dashboard: 1 },
        params: {},
        exp: Math.round(Date.now() / 1000) + (10 * 60) // 10 minute expiration
    },
}, ];