import{d as R,e as r,o as l,k as _,w as a,a as o,i as u,j as w,aq as E,P as C,b as n,t as i,S as b,$ as B,c as m,F as h,R as I,ah as L,A as N,l as f,C as X,m as q}from"./index-loxRIpgb.js";import{F}from"./FilterBar-D758saLR.js";import{S as T}from"./SummaryView-Bl9n50Fq.js";const $={class:"stack"},j={class:"columns"},G={key:0},K={key:1},O=R({__name:"ServiceDetailView",setup(W){return(Z,H)=>{const y=r("DataLoader"),v=r("KCard"),g=r("XAction"),x=r("XIcon"),V=r("XActionGroup"),S=r("RouterView"),D=r("DataCollection"),A=r("AppView"),P=r("RouteView");return l(),_(P,{name:"service-detail-view",params:{mesh:"",service:"",page:1,size:50,s:"",dataPlane:"",codeSearch:"",codeFilter:!1,codeRegExp:!1}},{default:a(({can:k,route:s,t:c,me:p,uri:z})=>[o(A,null,{default:a(()=>[u("div",$,[o(v,null,{default:a(()=>[o(y,{src:z(w(E),"/meshes/:mesh/service-insights/:name",{mesh:s.params.mesh,name:s.params.service})},{default:a(({data:t})=>{var e,d;return[u("div",j,[o(C,null,{title:a(()=>[n(i(c("http.api.property.status")),1)]),body:a(()=>[o(b,{status:t.status},null,8,["status"])]),_:2},1024),n(),o(C,null,{title:a(()=>[n(i(c("http.api.property.address")),1)]),body:a(()=>[t.addressPort?(l(),_(B,{key:0,text:t.addressPort},null,8,["text"])):(l(),m(h,{key:1},[n(i(c("common.detail.none")),1)],64))]),_:2},1024),n(),o(I,{online:((e=t.dataplanes)==null?void 0:e.online)??0,total:((d=t.dataplanes)==null?void 0:d.total)??0},{title:a(()=>[n(i(c("http.api.property.dataPlaneProxies")),1)]),_:2},1032,["online","total"])])]}),_:2},1032,["src"])]),_:2},1024),n(),u("div",null,[u("h3",null,i(c("services.detail.data_plane_proxies")),1),n(),o(v,{class:"mt-4"},{default:a(()=>[u("search",null,[o(F,{class:"data-plane-proxy-filter",placeholder:"name:dataplane-name",query:s.params.s,fields:{name:{description:"filter by name or parts of a name"},protocol:{description:"filter by “kuma.io/protocol” value"},tag:{description:"filter by tags (e.g. “tag: version:2”)"},...k("use zones")&&{zone:{description:"filter by “kuma.io/zone” value"}}},onChange:t=>s.update({...Object.fromEntries(t.entries())})},null,8,["query","fields","onChange"])]),n(),o(y,{src:z(w(L),"/meshes/:mesh/dataplanes/for/service-insight/:service",{mesh:s.params.mesh,service:s.params.service},{page:s.params.page,size:s.params.size,search:s.params.s})},{loadable:a(({data:t})=>[o(D,{type:"data-planes",items:(t==null?void 0:t.items)??[void 0],page:s.params.page,"page-size":s.params.size,total:t==null?void 0:t.total,onChange:s.update},{default:a(()=>[o(N,{class:"data-plane-collection","data-testid":"data-plane-collection",headers:[{...p.get("headers.name"),label:"Name",key:"name"},{...p.get("headers.namespace"),label:"Namespace",key:"namespace"},...k("use zones")?[{...p.get("headers.zone"),label:"Zone",key:"zone"}]:[],{...p.get("headers.certificate"),label:"Certificate Info",key:"certificate"},{...p.get("headers.status"),label:"Status",key:"status"},{...p.get("headers.warnings"),label:"Warnings",key:"warnings",hideLabel:!0},{...p.get("headers.actions"),label:"Actions",key:"actions",hideLabel:!0}],items:t==null?void 0:t.items,"is-selected-row":e=>e.name===s.params.dataPlane,onResize:p.set},{name:a(({row:e})=>[o(g,{"data-action":"",class:"name-link",to:{name:"service-data-plane-summary-view",params:{mesh:e.mesh,dataPlane:e.id},query:{page:s.params.page,size:s.params.size,s:s.params.s}}},{default:a(()=>[n(i(e.name),1)]),_:2},1032,["to"])]),namespace:a(({row:e})=>[n(i(e.namespace),1)]),zone:a(({row:e})=>[e.zone?(l(),_(g,{key:0,to:{name:"zone-cp-detail-view",params:{zone:e.zone}}},{default:a(()=>[n(i(e.zone),1)]),_:2},1032,["to"])):(l(),m(h,{key:1},[n(i(c("common.collection.none")),1)],64))]),certificate:a(({row:e})=>{var d;return[(d=e.dataplaneInsight.mTLS)!=null&&d.certificateExpirationTime?(l(),m(h,{key:0},[n(i(c("common.formats.datetime",{value:Date.parse(e.dataplaneInsight.mTLS.certificateExpirationTime)})),1)],64)):(l(),m(h,{key:1},[n(i(c("data-planes.components.data-plane-list.certificate.none")),1)],64))]}),status:a(({row:e})=>[o(b,{status:e.status},null,8,["status"])]),warnings:a(({row:e})=>[e.isCertExpired||e.warnings.length>0?(l(),_(x,{key:0,class:"mr-1",name:"warning"},{default:a(()=>[u("ul",null,[e.warnings.length>0?(l(),m("li",G,i(c("data-planes.components.data-plane-list.version_mismatch")),1)):f("",!0),n(),e.isCertExpired?(l(),m("li",K,i(c("data-planes.components.data-plane-list.cert_expired")),1)):f("",!0)])]),_:2},1024)):(l(),m(h,{key:1},[n(i(c("common.collection.none")),1)],64))]),actions:a(({row:e})=>[o(V,null,{default:a(()=>[o(g,{to:{name:"data-plane-detail-view",params:{dataPlane:e.id}}},{default:a(()=>[n(i(c("common.collection.actions.view")),1)]),_:2},1032,["to"])]),_:2},1024)]),_:2},1032,["headers","items","is-selected-row","onResize"]),n(),o(S,null,{default:a(({Component:e})=>[s.child()?(l(),_(T,{key:0,onClose:d=>s.replace({name:s.name,params:{mesh:s.params.mesh},query:{page:s.params.page,size:s.params.size,s:s.params.s}})},{default:a(()=>[typeof t<"u"?(l(),_(X(e),{key:0,items:t.items},null,8,["items"])):f("",!0)]),_:2},1032,["onClose"])):f("",!0)]),_:2},1024)]),_:2},1032,["items","page","page-size","total","onChange"])]),_:2},1032,["src"])]),_:2},1024)])])]),_:2},1024)]),_:1})}}}),U=q(O,[["__scopeId","data-v-6253b442"]]);export{U as default};
