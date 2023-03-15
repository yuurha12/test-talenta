import axios from 'axios';
import { downloadPdf } from './App';

jest.mock('axios');

test('downloadPdf should download PDF Report', async () => {
  const blob = new Blob(['test']);
  const responseObj = {data: blob};
  axios.get.mockResolvedValue(responseObj);

  document.body.appendChild = jest.fn();
  URL.createObjectURL = jest.fn(() => 'http://localhost:5000/api/v1');
  const link = { setAttribute: jest.fn(), click: jest.fn(), remove: jest.fn() }
  global.URL = { revokeObjectURL: jest.fn() }
  global.document.createElement = jest.fn(() => link)

  await downloadPdf();

  expect(axios.get).toHaveBeenCalledWith('/report', { responseType: 'blob' });
  expect(URL.createObjectURL).toHaveBeenCalledWith(blob);
  expect(link.setAttribute).toHaveBeenCalledWith('download', 'report.pdf');
  expect(document.body.appendChild).toHaveBeenCalledWith(link);
  expect(link.click).toHaveBeenCalled();
  expect(link.remove).toHaveBeenCalled();
});
