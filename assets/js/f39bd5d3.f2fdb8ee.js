"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[5321],{39013:(n,e,i)=>{i.r(e),i.d(e,{assets:()=>y,contentTitle:()=>m,default:()=>C,frontMatter:()=>k,metadata:()=>t,toc:()=>g});const t=JSON.parse('{"id":"1backend/container-is-running","title":"Check If a Container Is Running","description":"Check if a Docker container is running, identified by hash or name.","source":"@site/docs/1backend/container-is-running.api.mdx","sourceDirName":"1backend","slug":"/1backend/container-is-running","permalink":"/docs/1backend/container-is-running","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"container-is-running","title":"Check If a Container Is Running","description":"Check if a Docker container is running, identified by hash or name.","sidebar_label":"Check If a Container Is Running","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VU2P2zYQ/SsEz7Jk1y0K6NRNG6Ru0uyizp42PtDUWGJMkcyQsusK+u/BiLKldbw5BEUuliXOx5vHeTMtL8BLVC4oa3jOf69A7pnaMcH+sHIPyKQ1QSgDyJRn2BijTJkwVYAJaqegYNsTq4SvmEVmRA0pT7h1gIIirgqe80uElf8n+vOEO4GihgDoef7UXqO45PxT+IonXNHXzw3giSecsvCcV/HIywpqwfOWh5Oj7z4gpei65OWw7ynE7bAmHr0cdpNwBO+s8eDp/Kf5nB7PU92/5UlfOJhAp8I5rWTPSfbJk0k7SeGQGAsqBlQXmsbsW2s1CENVcYTPjUIoeP40sd0kZ1u7/QQy8K4j459voVuZg9CqYH+t79/Tvf2tvFemZA/jpXw3ekC0ePM6XsC3+BrfoxFNqCyq/6D4YUh+uc1UADRCszXgAZC97mP+GEhdwj3IBlU49Rp5BQIB75pQ8fxpQ30YREnymTT2+iCpE2oIlSXtlRB6sZEPzy5KnPmDHN8y5Wd4UabvK42ybFCTX6atFLqyPuSLxXL5K6fkZ2xrqjWWN0V4zeSHkwP2cTD5yNnOam2PcX4I5p2QwIQpWLB7MEzI2ONsh7ZmoQL26GN57J0tlWFgCmeVCelZxhWIAnDU8d3QQf3F8Au/wqm3cOoZV2ZnCWfPg+zvEmqhqGIvNPjfSBONFgGtSaWtJ7EfVmzdOGeR2I0kVSG4PMsWWyH3YApyyHiXXLFwt5oZEdQBWK0kWuJaSfDMaRF2FmsqRysJxgPhOed78/COHZbp/Fk2n2fZ8XhMS9OkFsts8POZKJ2eLdN5WoVaE4YAWPv73TpmG8H6oyhLwFTZrDfJiCcVNJksXsVCeMKpHSL8ebpM5zOUKfVAwp31oRZmgjTujxXtj7EpV56Ng/8ZH+0opP9p9QzXHODfkDktlCGcPWftIIMn/kwGg5j7d+qlUQqbhFPLk0fbboWHR9RdR5/jyuj3lvJiq2kW74T28K3yrnfaTaR7OI2r7SB0Qya8X2XfkWlYc9/INGy7MdOGXlBRqn7KJGdhUbXR505KcGHi9dUE7KYj6M3rDzzhYhgKowwpWHL+Q9HPR+Y0iX0t4wiBfomUmy5tG0XedRf7ePSix2V2RGsiaNN13Rcw8h8k","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Run a Container","permalink":"/docs/1backend/run-container"},"next":{"title":"Stop a Container","permalink":"/docs/1backend/stop-container"}}');var s=i(74848),a=i(28453),o=i(53746),r=i.n(o),c=i(56518),p=i.n(c),d=i(99972),l=i.n(d),h=i(25342),u=i.n(h),f=(i(44215),i(82223),i(24861));const k={id:"container-is-running",title:"Check If a Container Is Running",description:"Check if a Docker container is running, identified by hash or name.",sidebar_label:"Check If a Container Is Running",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VU2P2zYQ/SsEz7Jk1y0K6NRNG6Ru0uyizp42PtDUWGJMkcyQsusK+u/BiLKldbw5BEUuliXOx5vHeTMtL8BLVC4oa3jOf69A7pnaMcH+sHIPyKQ1QSgDyJRn2BijTJkwVYAJaqegYNsTq4SvmEVmRA0pT7h1gIIirgqe80uElf8n+vOEO4GihgDoef7UXqO45PxT+IonXNHXzw3giSecsvCcV/HIywpqwfOWh5Oj7z4gpei65OWw7ynE7bAmHr0cdpNwBO+s8eDp/Kf5nB7PU92/5UlfOJhAp8I5rWTPSfbJk0k7SeGQGAsqBlQXmsbsW2s1CENVcYTPjUIoeP40sd0kZ1u7/QQy8K4j459voVuZg9CqYH+t79/Tvf2tvFemZA/jpXw3ekC0ePM6XsC3+BrfoxFNqCyq/6D4YUh+uc1UADRCszXgAZC97mP+GEhdwj3IBlU49Rp5BQIB75pQ8fxpQ30YREnymTT2+iCpE2oIlSXtlRB6sZEPzy5KnPmDHN8y5Wd4UabvK42ybFCTX6atFLqyPuSLxXL5K6fkZ2xrqjWWN0V4zeSHkwP2cTD5yNnOam2PcX4I5p2QwIQpWLB7MEzI2ONsh7ZmoQL26GN57J0tlWFgCmeVCelZxhWIAnDU8d3QQf3F8Au/wqm3cOoZV2ZnCWfPg+zvEmqhqGIvNPjfSBONFgGtSaWtJ7EfVmzdOGeR2I0kVSG4PMsWWyH3YApyyHiXXLFwt5oZEdQBWK0kWuJaSfDMaRF2FmsqRysJxgPhOed78/COHZbp/Fk2n2fZ8XhMS9OkFsts8POZKJ2eLdN5WoVaE4YAWPv73TpmG8H6oyhLwFTZrDfJiCcVNJksXsVCeMKpHSL8ebpM5zOUKfVAwp31oRZmgjTujxXtj7EpV56Ng/8ZH+0opP9p9QzXHODfkDktlCGcPWftIIMn/kwGg5j7d+qlUQqbhFPLk0fbboWHR9RdR5/jyuj3lvJiq2kW74T28K3yrnfaTaR7OI2r7SB0Qya8X2XfkWlYc9/INGy7MdOGXlBRqn7KJGdhUbXR505KcGHi9dUE7KYj6M3rDzzhYhgKowwpWHL+Q9HPR+Y0iX0t4wiBfomUmy5tG0XedRf7ePSix2V2RGsiaNN13Rcw8h8k",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},m=void 0,y={},g=[];function b(n){const e={p:"p",...(0,a.R)(),...n.components};return(0,s.jsxs)(s.Fragment,{children:[(0,s.jsx)(f.default,{as:"h1",className:"openapi__heading",children:"Check If a Container Is Running"}),"\n",(0,s.jsx)(r(),{method:"get",path:"/container-svc/container/is-running",context:"endpoint"}),"\n",(0,s.jsx)(e.p,{children:"Check if a Docker container is running, identified by hash or name."}),"\n",(0,s.jsx)(f.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,s.jsx)(p(),{parameters:[{description:"Container Hash",in:"query",name:"hash",schema:{type:"string"}},{description:"Container Name",in:"query",name:"name",schema:{type:"string"}}]}),"\n",(0,s.jsx)(l(),{title:"Body",body:void 0}),"\n",(0,s.jsx)(u(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{properties:{isRunning:{type:"boolean"}},required:["isRunning"],type:"object"}}}},400:{description:"Invalid JSON or Missing Parameters",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function C(n={}){const{wrapper:e}={...(0,a.R)(),...n.components};return e?(0,s.jsx)(e,{...n,children:(0,s.jsx)(b,{...n})}):b(n)}}}]);