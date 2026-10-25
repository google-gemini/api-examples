
/**
  * @license
  * Copyright 2025 Google LLC
  *
  * Licensed under the Apache License, Version 2.0 (the "License");
  * you may not use this file except in compliance with the License.
  * You may obtain a copy of the License at
  *
  *   http://www.apache.org/licenses/LICENSE-2.0
  *
  * Unless required by applicable law or agreed to in writing, software
  * distributed under the License is distributed on an "AS IS" BASIS,
  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  * See the License for the specific language governing permissions and
  * limitations under the License.
  */

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;
import com.google.genai.Client;
import java.io.IOException;
import org.apache.http.HttpException;

/**
 * Tests for the CodeExecution class.
 * These tests verify the functionality of the code execution examples.
 */
public class CodeExecutionTest {

    @Test
    public void testCodeExecutionBasic() throws IOException, HttpException {
        Client client = new Client();
        String result = CodeExecution.codeExecutionBasic(client);

        // Check that the response contains non-empty text
        assertNotNull(result);
        assertTrue(result.length() > 0);
    }

    @Test
    public void testCodeExecutionRequestOverride() throws IOException, HttpException {
        Client client = new Client();
        String[] result = CodeExecution.codeExecutionRequestOverride(client);

        // Check that the response contains non-empty executable code and result
        assertNotNull(result);
        assertEquals(2, result.length);
        assertTrue(result[0].length() > 0); // executableCode
        assertTrue(result[1].length() > 0); // executionResult
    }
}
