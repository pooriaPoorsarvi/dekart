import { GetEnvRequest, GetEnvResponse } from '../../proto/dekart_pb'
import { Dekart } from '../../proto/dekart_pb_service'
import { unary } from '../lib/grpc'
import { error } from './message'
import * as Sentry from "@sentry/react";
// this creates the pay load
export function setEnv (variables) {
  return { type: setEnv.name, variables }
}

const typeToName = Object.keys(GetEnvResponse.Variable.Type).map(n => n.slice(5))

export function getEnv () {
  return async dispatch => {
    dispatch({ type: getEnv.name })
    const req = new GetEnvRequest()
    try {
      const { variablesList } = await unary(Dekart.GetEnv, req)
      const variables = variablesList.reduce((variables, v) => {
        variables[typeToName[v.type]] = v.value
        return variables
      }, {})
      dispatch(setEnv(variables))
      if(variables["SENTRY_DSN_FRONTEND"] && variables["SENTRY_DSN_FRONTEND"].length > 0){
        Sentry.init({
          dsn: variables["SENTRY_DSN_FRONTEND"],
          integrations: [new Sentry.BrowserTracing({ tracingOrigins: ["*"] })],

          // We recommend adjusting this value in production, or using tracesSampler
          // for finer control
          release: "1.41",
          tracesSampleRate: 1.0,
        });
        console.log("Sentry setup finished");
      }
    } catch (err) {
      dispatch(error(err))
    }
  }
}
