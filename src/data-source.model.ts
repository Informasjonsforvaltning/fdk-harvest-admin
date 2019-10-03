import { Document, model, Schema } from 'mongoose';
import pick from 'lodash/pick';

interface DataSource {
  // set to any because of mongoose type compatibility, in reality it is string
  id?: any; // eslint-disable-line @typescript-eslint/no-explicit-any
  dataSourceType?: string;
  url?: string;
  publisherId?: string;
  description?: string;
}

export interface DataSourceDocument extends Document, DataSource {}

const dataSourceSchemaDefinition = {
  id: String,
  dataSourceType: String,
  url: String,
  publisherId: String,
  description: String
};

const dataSourceFromDocument = (document: DataSourceDocument): DataSource =>
  pick(document, Object.getOwnPropertyNames(dataSourceSchemaDefinition));

const DataSourceSchema = new Schema(dataSourceSchemaDefinition, {
  timestamps: true,
  strict: true,
  toObject: {
    transform: dataSourceFromDocument
  }
});

export const DataSourceModel = model<DataSourceDocument>(
  'datasource',
  DataSourceSchema
);
