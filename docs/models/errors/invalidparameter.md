# InvalidParameter

Details why a certain parameter is invalid.


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `Detail`                                                        | **string*                                                       | :heavy_minus_sign:                                              | A human-readable explanation.                                   |
| `In`                                                            | [sdkerrors.In](../../models/errors/in.md)                       | :heavy_check_mark:                                              | N/A                                                             |
| `Name`                                                          | **string*                                                       | :heavy_minus_sign:                                              | The name of the invalid parameter                               |
| `Path`                                                          | **string*                                                       | :heavy_minus_sign:                                              | The JSON Path in the schema that the invalid parameter violated |
| `Title`                                                         | *string*                                                        | :heavy_check_mark:                                              | The reason this parameter is invalid                            |
| `Type`                                                          | *string*                                                        | :heavy_check_mark:                                              | A URI reference that identifies the problem type.               |