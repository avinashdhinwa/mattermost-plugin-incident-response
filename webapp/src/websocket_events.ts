// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import {GetStateFunc} from 'mattermost-redux/types/actions';
import {WebSocketMessage} from 'mattermost-redux/actions/websocket';
import {getCurrentTeam} from 'mattermost-redux/selectors/entities/teams';

import {Dispatch} from 'redux';

import {navigateToUrl} from 'src/browser_routing';
import {incidentCreated, receivedTeamIncidentChannels} from 'src/actions';
import {fetchIncidentChannels} from 'src/client';

import {clientId} from './selectors';
import {Incident, isIncident} from './types/incident';

export const websocketSubscribers = new Set<(incident: Incident) => void>();

export function handleReconnect(getState: GetStateFunc, dispatch: Dispatch) {
    return async (): Promise<void> => {
        const currentTeam = getCurrentTeam(getState());
        dispatch(receivedTeamIncidentChannels(await fetchIncidentChannels(currentTeam.id)));
    };
}

export function handleWebsocketIncidentUpdate() {
    return (msg: WebSocketMessage): void => {
        if (!msg.data.payload) {
            return;
        }
        const data = JSON.parse(msg.data.payload);

        // eslint-disable-next-line no-process-env
        if (process.env.NODE_ENV !== 'production') {
            if (!isIncident(data)) {
                // eslint-disable-next-line no-console
                console.error('received a websocket data payload that was not an incident in handleWebsocketIncidentUpdate:', data);
            }
        }
        const incident = data as Incident;

        websocketSubscribers.forEach((fn) => fn(incident));
    };
}

export function handleWebsocketIncidentCreate(getState: GetStateFunc, dispatch: Dispatch) {
    return (msg: WebSocketMessage): void => {
        if (!msg.data.payload) {
            return;
        }
        const payload = JSON.parse(msg.data.payload);
        const data = payload.incident;

        // eslint-disable-next-line no-process-env
        if (process.env.NODE_ENV !== 'production') {
            if (!isIncident(data)) {
                // eslint-disable-next-line no-console
                console.error('received a websocket data payload that was not an incident in handleWebsocketIncidentCreate:', data);
            }
        }
        const incident = data as Incident;

        dispatch(incidentCreated(incident));

        if (payload.client_id !== clientId(getState())) {
            return;
        }

        const currentTeam = getCurrentTeam(getState());

        // Navigate to the newly created channel
        const url = `/${currentTeam.name}/channels/${incident.channel_id}`;
        navigateToUrl(url);
    };
}
