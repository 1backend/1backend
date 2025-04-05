"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[6756],{25407:(e,t,i)=>{i.r(t),i.d(t,{assets:()=>b,contentTitle:()=>j,default:()=>X,frontMatter:()=>g,metadata:()=>o,toc:()=>y});const o=JSON.parse('{"id":"1backend/upload-file","title":"Upload a File","description":"Uploads a file to the server.","source":"@site/docs/1backend/upload-file.api.mdx","sourceDirName":"1backend","slug":"/1backend/upload-file","permalink":"/docs/1backend/upload-file","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"upload-file","title":"Upload a File","description":"Uploads a file to the server.","sidebar_label":"Upload a File","hide_title":true,"hide_table_of_contents":true,"api":"eJztVk1vGzcQ/SsDnlpgtWtVbgvoVCdpCqFJbcR2gcIyYIo7u8uYS9L8kKMI+9+LIVe2bCtB0UPQQy/Sgh9v3gxn3syW1eiFkzZIo9mcXVpleO2BQyMVQjAQOgSPbo2uXOrX0TnUQW1ANhC91G3aF0qiDh6MVhswGvNlwTWsEGKCxBp4AA5B9gh1HKGlh1VsoQvB+nlVtTJ0cVUK01enFvXJ2eLCGOUrY1FzKyctanQ8GFdJ7yP6ajqdHU+X+lQLhNDxANJDIz9hDfcIvjNR1dDxNQLXcHN2eQEVMZv4tagyLX8zfryVCj18p03A5JJV0XH1PaCurZE6+HKpFxocciXDJh3ZbcGSRW2aRgrJldosGfhorXHBQx9VkFblgPgS/nr//s9yqZf6A95F6dAnoJsdqXnmMhcOecAbsOh66b00umQFM5acl0YvajZnj7RZwRzeRfThlak3bL5lwuiAOtBnZsBdqBrj+knNA6dlLzrs01fYWGRzZlYfUQRWMOvITpDoaZeY0f/TLHk75kbmwIodiA9O6pYVjGzxwOZsJTV3GzYMmaN0WLP5VYa9Hp4uBxcxLXhrtM/2fzg6+oL5h7TyUQj0volKbVix7zu3VkmRQlZ99HR13/Gnjo6uvFjPT1GfhL1YjW4ORfJjUb9k+M60UnCV62DxBrzlWlO1cKXAYaLlX8RtBPyD91+IueY9UopT0hgnW6m5grRomrTY5HQ4CHvGQ3cYlnZ2sBRG0KZGUIY8sLT3D9DP5edE+uHlpQ4/HT+eljpgi46OywMBu9TyLqZQNcZlYRjDdMgg8TsU9sWbHdXkgQ/GZYn6Oly09VfeOHp02dizrQM5naJwXTwrKjr5fIXWjg8l90KvuZI1jCX97zManTPuMO3DZKaHnoXH0BknP2P9zZj8eDgsAR2l+3nqRfBrwvw2lIaCeRTRybBh86ste4XcoTuJVE9X1wO9N2895UBSpvO1oBzoMXSGlNrGpKup/NjzBsQImzzyCTo6RYeqVHyd8WE+nc5mPzMysuNwTj5lN/aZPI/YxcYiLMcjSwaNUcrcYw2rDXBSJEFtsYZgblEDFzmToXGmTyV06dGRL0Biph+6HXUiSfgd8hrpCXTSK3YyZkp6gMcy41b+jrkDSN2YXXviIr0Z9lySx54r9L/QQBEVD85oGgL2sM8WcJ6bKivGINHQMK+q6YqLW9R1mhqoXp9G4WQx0TzINUIvhTMUaynQg1U8kFaRO0oK1D7J187eb2fvYD0rj55YoxHl/v6+bHUsjWur8Z6veGvVZFYelV3oFXEI6Hp/2pxna49k/T1vW3SlNFU6UlGcZKAey6avsiOsYJQOmf5ROSuPJk6Us2PCtcaHnus9pnlgAw7jJPDE+71J4P/J7j882Y2lEvBTqKziUqe+Q3m3HXUjtxdCoozMynFdMFII2txuV9zjpVPDQMt3ER1J1XXB1txJvqIEu7oeil3RktTc4obN2eucIBNSC8o8riJROTQ2DsXu0okQaMPe8RfSO+wL4NnlBSvYahxOe1PTFcJNsMXjZ+bIRzl7FBCyWuw+iPtuS2/2SDwXoMyVfon5wSvbbZanYXg4n7e+eONB9fJpeiqaY/8G7f6ftQ==","sidebar_class_name":"put api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Serve an Uploaded File","permalink":"/docs/1backend/serve-upload"},"next":{"title":"List Uploads","permalink":"/docs/1backend/list-uploads"}}');var n=i(74848),l=i(28453),s=i(53746),a=i.n(s),r=i(56518),p=i.n(r),d=i(99972),c=i.n(d),u=i(25342),h=i.n(u),f=(i(44215),i(82223),i(24861));const g={id:"upload-file",title:"Upload a File",description:"Uploads a file to the server.",sidebar_label:"Upload a File",hide_title:!0,hide_table_of_contents:!0,api:"eJztVk1vGzcQ/SsDnlpgtWtVbgvoVCdpCqFJbcR2gcIyYIo7u8uYS9L8kKMI+9+LIVe2bCtB0UPQQy/Sgh9v3gxn3syW1eiFkzZIo9mcXVpleO2BQyMVQjAQOgSPbo2uXOrX0TnUQW1ANhC91G3aF0qiDh6MVhswGvNlwTWsEGKCxBp4AA5B9gh1HKGlh1VsoQvB+nlVtTJ0cVUK01enFvXJ2eLCGOUrY1FzKyctanQ8GFdJ7yP6ajqdHU+X+lQLhNDxANJDIz9hDfcIvjNR1dDxNQLXcHN2eQEVMZv4tagyLX8zfryVCj18p03A5JJV0XH1PaCurZE6+HKpFxocciXDJh3ZbcGSRW2aRgrJldosGfhorXHBQx9VkFblgPgS/nr//s9yqZf6A95F6dAnoJsdqXnmMhcOecAbsOh66b00umQFM5acl0YvajZnj7RZwRzeRfThlak3bL5lwuiAOtBnZsBdqBrj+knNA6dlLzrs01fYWGRzZlYfUQRWMOvITpDoaZeY0f/TLHk75kbmwIodiA9O6pYVjGzxwOZsJTV3GzYMmaN0WLP5VYa9Hp4uBxcxLXhrtM/2fzg6+oL5h7TyUQj0volKbVix7zu3VkmRQlZ99HR13/Gnjo6uvFjPT1GfhL1YjW4ORfJjUb9k+M60UnCV62DxBrzlWlO1cKXAYaLlX8RtBPyD91+IueY9UopT0hgnW6m5grRomrTY5HQ4CHvGQ3cYlnZ2sBRG0KZGUIY8sLT3D9DP5edE+uHlpQ4/HT+eljpgi46OywMBu9TyLqZQNcZlYRjDdMgg8TsU9sWbHdXkgQ/GZYn6Oly09VfeOHp02dizrQM5naJwXTwrKjr5fIXWjg8l90KvuZI1jCX97zManTPuMO3DZKaHnoXH0BknP2P9zZj8eDgsAR2l+3nqRfBrwvw2lIaCeRTRybBh86ste4XcoTuJVE9X1wO9N2895UBSpvO1oBzoMXSGlNrGpKup/NjzBsQImzzyCTo6RYeqVHyd8WE+nc5mPzMysuNwTj5lN/aZPI/YxcYiLMcjSwaNUcrcYw2rDXBSJEFtsYZgblEDFzmToXGmTyV06dGRL0Biph+6HXUiSfgd8hrpCXTSK3YyZkp6gMcy41b+jrkDSN2YXXviIr0Z9lySx54r9L/QQBEVD85oGgL2sM8WcJ6bKivGINHQMK+q6YqLW9R1mhqoXp9G4WQx0TzINUIvhTMUaynQg1U8kFaRO0oK1D7J187eb2fvYD0rj55YoxHl/v6+bHUsjWur8Z6veGvVZFYelV3oFXEI6Hp/2pxna49k/T1vW3SlNFU6UlGcZKAey6avsiOsYJQOmf5ROSuPJk6Us2PCtcaHnus9pnlgAw7jJPDE+71J4P/J7j882Y2lEvBTqKziUqe+Q3m3HXUjtxdCoozMynFdMFII2txuV9zjpVPDQMt3ER1J1XXB1txJvqIEu7oeil3RktTc4obN2eucIBNSC8o8riJROTQ2DsXu0okQaMPe8RfSO+wL4NnlBSvYahxOe1PTFcJNsMXjZ+bIRzl7FBCyWuw+iPtuS2/2SDwXoMyVfon5wSvbbZanYXg4n7e+eONB9fJpeiqaY/8G7f6ftQ==",sidebar_class_name:"put api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},j=void 0,b={},y=[];function v(e){const t={a:"a",code:"code",p:"p",...(0,l.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(f.default,{as:"h1",className:"openapi__heading",children:"Upload a File"}),"\n",(0,n.jsx)(a(),{method:"put",path:"/file-svc/upload",context:"endpoint"}),"\n",(0,n.jsxs)(t.p,{children:["Uploads a file to the server.\nCurrently if using the clients only one file can be uploaded at a time due to this bug ",(0,n.jsx)(t.a,{href:"https://github.com/OpenAPITools/openapi-generator/issues/11341",children:"https://github.com/OpenAPITools/openapi-generator/issues/11341"}),"\nOnce that is fixed we should have an ",(0,n.jsx)(t.code,{children:"PUT /file-svc/uploads"}),'/uploadFiles (note the plural) endpoints.\nIn reality the endpoint "unofficially" supports multiple files. YMMV.']}),"\n",(0,n.jsxs)(t.p,{children:["Requires the ",(0,n.jsx)(t.code,{children:"file-svc:upload:create"})," permission."]}),"\n",(0,n.jsx)(f.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,n.jsx)(p(),{parameters:void 0}),"\n",(0,n.jsx)(c(),{title:"Body",body:{content:{"multipart/form-data":{schema:{type:"object",properties:{file:{description:"File to upload",type:"string",format:"binary"}},required:["file"]}}},required:!0}}),"\n",(0,n.jsx)(h(),{id:void 0,label:void 0,responses:{200:{description:"File uploaded successfully",content:{"application/json":{schema:{properties:{upload:{properties:{createdAt:{type:"string"},fileId:{description:"Logical file ID spanning all replicas",type:"string"},fileName:{description:"Filename is the original name of the file",type:"string"},filePath:{description:"FilePath is the full node local path of the file",type:"string"},fileSize:{format:"int64",type:"integer"},id:{description:"Unique ID for this replica",type:"string"},nodeId:{description:"ID of the node storing this replica",type:"string"},updatedAt:{type:"string"},userId:{type:"string"}},required:["fileSize"],type:"object"}},type:"object"}}}},400:{description:"Invalid request",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function X(e={}){const{wrapper:t}={...(0,l.R)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(v,{...e})}):v(e)}}}]);