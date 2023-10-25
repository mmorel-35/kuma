import{d as O,g as z,f as L,h as v,r as P,o as n,i as y,w as e,j as l,a4 as S,n as t,p as r,k as p,a5 as C,a6 as _,H as s,a7 as H,a8 as K,K as F,m as N,l as u,F as m,I as b,a9 as $,t as j}from"./index-eb555afc.js";import{S as U}from"./StatusBadge-0c1ce262.js";import{_ as W}from"./SubscriptionList.vue_vue_type_script_setup_true_lang-49226749.js";import{T as R}from"./TagList-f6df787a.js";import{a as Z,d as x,b as G,c as X,C as q,I as J}from"./dataplane-0a086c06.js";import"./AccordionList-c93339c1.js";const Q=["data-testid","innerHTML"],Y={class:"stack","data-testid":"detail-view-details"},tt={class:"columns"},et={class:"status-with-reason"},at={class:"columns"},st=["innerHTML"],nt={key:0},it=O({__name:"DataPlaneDetailView",props:{data:{}},setup(B){const{formatIsoDate:w}=z(),M=L(),o=B,D=v(()=>Z(o.data.dataplane,o.data.dataplaneInsight)),T=v(()=>x(o.data.dataplane)),V=v(()=>G(o.data.dataplaneInsight)),E=v(()=>{var d,k;const h=((d=o.data.dataplaneInsight)==null?void 0:d.subscriptions)??[];if(h.length===0)return[];const f=h[h.length-1];if(!("version"in f)||!f.version)return[];const c=f.version,g=[];if(c.kumaDp&&c.envoy){const i=X(c);i.kind!==q&&i.kind!==J&&g.push(i)}const a=(k=o.data.dataplaneInsight)==null?void 0:k.mTLS;return a&&Date.now()>new Date(a==null?void 0:a.certificateExpirationTime).getTime()&&g.push({kind:"CERT_EXPIRED",payload:{}}),M("use zones")&&x(o.data.dataplane).find(A=>A.label==="kuma.io/zone")&&typeof c.kumaDp.kumaCpCompatible=="boolean"&&!c.kumaDp.kumaCpCompatible&&g.push({kind:"INCOMPATIBLE_ZONE_CP_AND_KUMA_DP_VERSIONS",payload:{kumaDp:c.kumaDp.version}}),g});return(h,f)=>{const c=P("AppView"),g=P("RouteView");return n(),y(g,{name:"data-plane-detail-view"},{default:e(({t:a})=>[l(c,null,S({default:e(()=>{var d,k;return[t(),r("div",Y,[l(p(C),null,{body:e(()=>[r("div",tt,[l(_,null,{title:e(()=>[t(s(a("http.api.property.status")),1)]),body:e(()=>[r("div",et,[l(U,{status:D.value.status},null,8,["status"]),t(),D.value.reason.length>0?(n(),y(p(H),{key:0,label:D.value.reason.join(", "),class:"reason-tooltip"},{default:e(()=>[l(p(K),{size:p(F),"hide-title":""},null,8,["size"])]),_:1},8,["label"])):N("",!0)])]),_:2},1024),t(),l(_,null,{title:e(()=>[t(s(a("http.api.property.tags")),1)]),body:e(()=>[T.value.length>0?(n(),y(R,{key:0,tags:T.value},null,8,["tags"])):(n(),u(m,{key:1},[t(s(a("common.detail.none")),1)],64))]),_:2},1024),t(),l(_,null,{title:e(()=>[t(s(a("http.api.property.dependencies")),1)]),body:e(()=>[V.value!==null?(n(),y(R,{key:0,tags:V.value},null,8,["tags"])):(n(),u(m,{key:1},[t(s(a("common.detail.none")),1)],64))]),_:2},1024)])]),_:2},1024),t(),r("div",null,[r("h2",null,s(a("data-planes.routes.item.mtls.title")),1),t(),(d=o.data.dataplaneInsight)!=null&&d.mTLS?(n(!0),u(m,{key:0},b([o.data.dataplaneInsight.mTLS],i=>(n(),y(p(C),{key:i,class:"mt-4"},{body:e(()=>[r("div",at,[l(_,null,{title:e(()=>[t(s(a("data-planes.routes.item.mtls.expiration_time.title")),1)]),body:e(()=>[t(s(p(w)(i.certificateExpirationTime)),1)]),_:2},1024),t(),l(_,null,{title:e(()=>[t(s(a("data-planes.routes.item.mtls.generation_time.title")),1)]),body:e(()=>[t(s(p(w)(i.lastCertificateRegeneration)),1)]),_:2},1024),t(),l(_,null,{title:e(()=>[t(s(a("data-planes.routes.item.mtls.regenerations.title")),1)]),body:e(()=>[t(s(a("common.formats.integer",{value:i.certificateRegenerations})),1)]),_:2},1024),t(),l(_,null,{title:e(()=>[t(s(a("data-planes.routes.item.mtls.issued_backend.title")),1)]),body:e(()=>[t(s(i.issuedBackend),1)]),_:2},1024),t(),l(_,null,{title:e(()=>[t(s(a("data-planes.routes.item.mtls.supported_backends.title")),1)]),body:e(()=>[r("ul",null,[(n(!0),u(m,null,b(i.supportedBackends,I=>(n(),u("li",{key:I},s(I),1))),128))])]),_:2},1024)])]),_:2},1024))),128)):(n(),y(p($),{key:1,class:"mt-4",appearance:"warning"},{alertMessage:e(()=>[r("div",{innerHTML:a("data-planes.routes.item.mtls.disabled")},null,8,st)]),_:2},1024))]),t(),(n(!0),u(m,null,b([((k=o.data.dataplaneInsight)==null?void 0:k.subscriptions)??[]],i=>(n(),u(m,{key:i},[i.length>0?(n(),u("div",nt,[r("h2",null,s(a("data-planes.routes.item.subscriptions.title")),1),t(),l(p(C),{class:"mt-4"},{body:e(()=>[l(W,{subscriptions:i},null,8,["subscriptions"])]),_:2},1024)])):N("",!0)],64))),128))])]}),_:2},[E.value.length>0?{name:"notifications",fn:e(()=>[r("ul",null,[(n(!0),u(m,null,b(E.value,d=>(n(),u("li",{key:d.kind,"data-testid":`warning-${d.kind}`,innerHTML:a(`common.warnings.${d.kind}`,d.payload)},null,8,Q))),128)),t()])]),key:"0"}:void 0]),1024)]),_:1})}}});const ct=j(it,[["__scopeId","data-v-79d4a52d"]]);export{ct as default};