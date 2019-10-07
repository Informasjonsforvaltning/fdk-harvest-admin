import { Document, model, Schema } from 'mongoose';
import pick from 'lodash/pick';
import uuidv4 from 'uuid/v4';

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
  id: {
    default: uuidv4,
    type: String,
    required: true,
    unique: true
  },
  dataSourceType: {
    type: String,
    required: true
  },
  url: {
    type: String,
    required: true
  },
  publisherId: {
    type: String,
    required: true
  },
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
