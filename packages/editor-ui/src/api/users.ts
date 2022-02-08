import { IPersonalizationSurveyAnswers, IRestApiContext, IUserResponse } from '@/Interface';
import { IDataObject } from 'n8n-workflow';
import { makeRestApiRequest } from './helpers';

export function loginCurrentUser(context: IRestApiContext): Promise<IUserResponse | null> {
	return makeRestApiRequest(context, 'GET', '/login');
}

export function getCurrentUser(context: IRestApiContext): Promise<IUserResponse | null> {
	return makeRestApiRequest(context, 'GET', '/me');
}

export function login(context: IRestApiContext, params: {email: string, password: string}): Promise<IUserResponse> {
	return makeRestApiRequest(context, 'POST', '/login', params);
}

export async function logout(context: IRestApiContext): Promise<void> {
	await makeRestApiRequest(context, 'POST', '/logout');
}

export function setupOwner(context: IRestApiContext, params: { firstName: string; lastName: string; email: string; password: string;}): Promise<IUserResponse> {
	return makeRestApiRequest(context, 'POST', '/owner/setup', params as unknown as IDataObject);
}

export function validateSignupToken(context: IRestApiContext, params: {inviterId: string; inviteeId: string}): Promise<{inviter: {firstName: string, lastName: string}}> {
	return makeRestApiRequest(context, 'GET', '/resolve-signup-token', params);
}

export function signup(context: IRestApiContext, params:  {inviterId: string; inviteeId: string; firstName: string; lastName: string; password: string}): Promise<IUserResponse> {
	return makeRestApiRequest(context, 'POST', '/user', params as unknown as IDataObject);
}

export function sendForgotPasswordEmail(context: IRestApiContext, params: {email: string}): Promise<void> {
	makeRestApiRequest(context, 'POST', '/forgot-password', params);
}

export function validatePasswordToken(context: IRestApiContext, params: {token: string, userId: string}): Promise<void> {
	makeRestApiRequest(context, 'GET', '/resolve-password-token', params);
}

export function changePassword(context: IRestApiContext, params: {token: string, password: string, userId: string}): Promise<void> {
	makeRestApiRequest(context, 'POST', '/change-password', params);
}

export function updateCurrentUser(context: IRestApiContext, params: {id: string, firstName: string, lastName: string, email: string}): Promise<IUserResponse> {
	return makeRestApiRequest(context, 'PATCH', `/me`, params as unknown as IDataObject);
}

export function updateCurrentUserPassword(context: IRestApiContext, params: {password: string}): Promise<void> {
	return makeRestApiRequest(context, 'PATCH', `/me/password`, params);
}

export function deleteUser(context: IRestApiContext, {id, transferId}: {id: string, transferId?: string}): Promise<void> {
	makeRestApiRequest(context, 'DELETE', `/users/${id}`, transferId ? { transferId } : {});
}

export function getUsers(context: IRestApiContext): Promise<IUserResponse[]> {
	return makeRestApiRequest(context, 'GET', '/users');
}

export function inviteUsers(context: IRestApiContext, params: Array<{email: string}>): Promise<Array<Partial<IUserResponse>>> {
	return makeRestApiRequest(context, 'POST', '/users', params as unknown as IDataObject);
}

export async function reinvite(context: IRestApiContext, {id}: {id: string}): Promise<void> {
	await makeRestApiRequest(context, 'POST', `/users/${id}/reinvite`);
}

export function submitPersonalizationSurvey(context: IRestApiContext, params: IPersonalizationSurveyAnswers): Promise<void> {
	makeRestApiRequest(context, 'POST', '/me/survey', params as unknown as IDataObject);
}
