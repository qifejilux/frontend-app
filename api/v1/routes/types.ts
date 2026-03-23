// types.ts

import { Action, State } from './store';

export enum Theme {
  LIGHT,
  DARK
}

export interface IProject {
  id: string;
  name: string;
  description: string;
  type: string;
  status: string;
  startDate: Date;
  endDate: Date;
  userId: string;
}

export interface IProjectState extends State {
  projects: IProject[];
  project: IProject | null;
  loading: boolean;
  error: string | null;
}

export interface IProjectAction extends Action {
  type: 'GET_PROJECTS' | 'GET_PROJECT' | 'CREATE_PROJECT' | 'UPDATE_PROJECT' | 'DELETE_PROJECT';
  payload?: any;
}